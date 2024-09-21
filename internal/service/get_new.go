package service

import (
	"errors"
	"fmt"
	"time"
)

func (s *Service) HandleGetNew() (string, error) {
	s.mu.Lock()
	for k, proxy := range s.activeProxies {
		if time.Now().Sub(proxy.CreatedAt) > s.config.ProxyTTL {
			delete(s.activeProxies, k)
		}
	}
	s.mu.Unlock()
	if len(s.activeProxies) >= s.config.MaxProxyServers {
		return "", errors.New("достигнуто максимальное количество активных прокси-серверов")
	}
	s.mu.Lock()
	var nProxy int64
	for i := s.config.ProxyServerPort; i < s.config.ProxyServerPort+s.config.MaxProxyServers; i++ {
		if _, ok := s.activeProxies[i]; !ok {
			s.activeProxies[i] = ProxyServer{
				ID:        "proxy_" + time.Now().Format("20060102150405"),
				CreatedAt: time.Now(),
			}
			nProxy = int64(i)
			break
		}
	}
	s.mu.Unlock()

	proxyData, err := s.repo.FetchExternalData()
	if err != nil {
		deleteProxy(s, nProxy)
		return "", err
	}

	newProxy, err := s.repo.CreateProxyServer(proxyData, nProxy)
	if err != nil {
		deleteProxy(s, nProxy)
		return "", err
	}

	proxyResponse := fmt.Sprintf("%v://%v:%v@%v:%v", newProxy.Type, newProxy.Login, newProxy.Pass, s.config.ProxyServerHost, newProxy.Port)

	return proxyResponse, nil
}

func deleteProxy(s *Service, proxyID int64) {
	s.mu.Lock()
	delete(s.activeProxies, int(proxyID))
	_ = s.repo.CleanProxy(proxyID)
	s.mu.Unlock()
}
