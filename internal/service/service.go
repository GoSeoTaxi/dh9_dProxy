package service

import (
	"sync"
	"time"

	"github.com/GoSeoTaxi/dh9_dProxy/internal/config"
	"github.com/GoSeoTaxi/dh9_dProxy/internal/repository"
)

type ProxyServer struct {
	ID        string    // Идентификатор прокси-сервера
	CreatedAt time.Time // Время создания прокси-сервера
}

type Service struct {
	repo          *repository.Repository
	config        *config.Config
	activeProxies map[int]ProxyServer
	mu            sync.Mutex
}

func NewService(r *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		repo:          r,
		config:        cfg,
		activeProxies: make(map[int]ProxyServer),
		mu:            sync.Mutex{},
	}
}
