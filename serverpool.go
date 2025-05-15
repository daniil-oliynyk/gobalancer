package main

import "github.com/daniil-oliynyk/gobalancer/config"

type IServerPool interface {
	AddBackend(Backend)
	GetBackends() []Backend
}

type ServerPool struct {
	backends []IBackend
}

func (s *ServerPool) AddBackend(b IBackend) {
	s.backends = append(s.backends, b)
}

func (s *ServerPool) GetBackends() []IBackend {
	return s.backends
}

func NewServerPool(c *config.Config) ServerPool {
	serverPool := ServerPool{}
	for _, b := range c.Backends {
		backend := NewBackend(b)
		serverPool.backends = append(serverPool.backends, backend)
	}

	return serverPool
}
