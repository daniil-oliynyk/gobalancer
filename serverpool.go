package main

import (
	"github.com/daniil-oliynyk/gobalancer/config"
	"github.com/daniil-oliynyk/gobalancer/interfaces"
)

type ServerPool struct {
	backends []interfaces.IBackend
	size     uint32
}

func (s *ServerPool) AddBackend(b interfaces.IBackend) {
	s.backends = append(s.backends, b)
	s.size++
}

func (s *ServerPool) GetBackends() []interfaces.IBackend {
	return s.backends
}
func (s *ServerPool) GetSize() uint32 {
	return s.size
}

func NewServerPool(c *config.Config) *ServerPool {
	serverPool := &ServerPool{}
	for _, b := range c.Backends {
		backend := NewBackend(b)
		serverPool.AddBackend(backend)
	}

	return serverPool
}
