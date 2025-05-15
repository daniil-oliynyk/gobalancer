package roundrobin

import (
	"github.com/daniil-oliynyk/gobalancer/interfaces"
	"github.com/daniil-oliynyk/gobalancer/utils"
	"github.com/valyala/fasthttp"
)

var logger = utils.GetLogger()

type RoundRobin struct {
}

func (r *RoundRobin) Serve() func(ctx *fasthttp.RequestCtx) {
	return nil
}

func NewRoundRobin(sp interfaces.IServerPool) interfaces.ILoadBalancer {
	logger.Debug().Msg("NewRoundRobin.created")
	return nil
}
