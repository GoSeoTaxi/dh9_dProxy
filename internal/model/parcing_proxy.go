package model

import (
	"fmt"
	"net/url"
	"strings"
)

type Proxy struct {
	Type  string // Тип прокси (например, socks5)
	Login string // Логин пользователя
	Pass  string // Пароль пользователя
	Addr  string // Адрес и порт
}

func ParseProxy(proxyStr string) (*Proxy, error) {

	proxyStr = strings.TrimSpace(proxyStr)
	proxyStr = strings.ToLower(proxyStr)
	parsedURL, err := url.Parse(proxyStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга прокси: %w", err)
	}

	login := ""
	pass := ""
	if parsedURL.User != nil {
		login = parsedURL.User.Username()
		pass, _ = parsedURL.User.Password()
	}

	proxy := &Proxy{
		Type:  parsedURL.Scheme,
		Login: login,
		Pass:  pass,
		Addr:  parsedURL.Host,
	}

	return proxy, nil
}
