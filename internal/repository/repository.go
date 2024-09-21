package repository

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-resty/resty/v2"

	"github.com/GoSeoTaxi/dh9_dProxy/internal/config"
	"github.com/GoSeoTaxi/dh9_dProxy/internal/model"
)

const lenRandomString = 8

type Repository struct {
	client          *resty.Client
	externalAPIURLs []string   // Массив URL из конфигурации
	gostV3Host      string     // Хост сервера Gost v3
	gostV3Port      string     // Порт сервера Gost v3
	gostV3Username  string     // Логин для Gost v3
	gostV3Password  string     // Пароль для Gost v3
	currentURLIdx   int        // Текущий индекс для раунд-робина
	mu              sync.Mutex // Для защиты индекса при конкурентном доступе
}

func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		client:          resty.New(),
		externalAPIURLs: cfg.ExternalAPIURLs,
		gostV3Host:      cfg.GostV3Host,
		gostV3Port:      cfg.GostV3Port,
		gostV3Username:  cfg.GostV3Username,
		gostV3Password:  cfg.GostV3Password,
		currentURLIdx:   0,
	}
}

func (r *Repository) FetchExternalData() (string, error) {
	if len(r.externalAPIURLs) == 0 {
		return "", nil // Или возвращаем ошибку, если нет доступных URL
	}

	r.mu.Lock()
	currentURL := r.externalAPIURLs[r.currentURLIdx]
	r.currentURLIdx = (r.currentURLIdx + 1) % len(r.externalAPIURLs) // Раунд-робин
	r.mu.Unlock()

	// Выполняем запрос
	resp, err := r.client.R().Get(currentURL)
	if err != nil {
		return "", err
	}

	if (resp.StatusCode()) != 200 {
		return "", errors.New("Try again later")
	}

	return string(resp.Body()), nil
}

func (r *Repository) CreateProxyServer(proxyData string, namePrefix int64) (*model.NewProxy, error) {

	err := r.deleteProxyChain(fmt.Sprintf("chain-%d", namePrefix))
	if err != nil {
		return nil, err
	}
	err = r.deleteService(fmt.Sprintf("service-%d", namePrefix))
	if err != nil {
		return nil, err
	}

	err = r.addChain(proxyData, namePrefix)
	if err != nil {
		return nil, err
	}

	var newProxy model.NewProxy
	newProxy.Type = "socks5"
	newProxy.Login = model.GenerateRandomString(lenRandomString)
	newProxy.Pass = model.GenerateRandomString(lenRandomString)
	newProxy.Port = fmt.Sprintf("%v", namePrefix)

	err = r.addService(fmt.Sprintf("service-%d", namePrefix), fmt.Sprintf("chain-%d", namePrefix),
		fmt.Sprintf("%v", newProxy.Port), newProxy.Type, newProxy.Login, newProxy.Pass)
	if err != nil {
		return nil, err
	}
	return &newProxy, nil
}

func (r *Repository) CleanProxy(namePrefix int64) error {
	_ = r.deleteProxyChain(fmt.Sprintf("chain-%d", namePrefix))
	_ = r.deleteService(fmt.Sprintf("service-%d", namePrefix))
	return nil
}
