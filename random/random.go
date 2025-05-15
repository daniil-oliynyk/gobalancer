package random

import (
	"math/rand/v2"

	"github.com/daniil-oliynyk/gobalancer/interfaces"
	"github.com/daniil-oliynyk/gobalancer/utils"
	"github.com/valyala/fasthttp"
)

var logger = utils.GetLogger()

type Random struct {
	serverPool interfaces.IServerPool
}

func (r *Random) Serve() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		b := r.next()
		b.ProxyHandler(ctx)
	}
}
func NewRandom(sp interfaces.IServerPool) interfaces.ILoadBalancer {
	logger.Debug().Msg("NewRandom.created")
	random := &Random{
		serverPool: sp,
	}

	return random
}

func (r *Random) next() interfaces.IBackend {
	spSize := r.serverPool.GetSize()
	return r.serverPool.GetBackends()[rand.Uint32N(spSize)]
}
