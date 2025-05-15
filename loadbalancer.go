package main

import (
	"github.com/daniil-oliynyk/gobalancer/config"
	"github.com/daniil-oliynyk/gobalancer/interfaces"
	Random "github.com/daniil-oliynyk/gobalancer/random"
)

func NewLoadBalancerHandler(c *config.Config, sp *ServerPool) interfaces.ILoadBalancer {

	switch c.Algo {
	case "random":
		return Random.NewRandom(sp)
	case "round-robin":

	}

	return nil
}
