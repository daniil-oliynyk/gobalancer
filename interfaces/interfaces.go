package interfaces

import "github.com/valyala/fasthttp"

type ILoadBalancer interface {
	Serve() func(ctx *fasthttp.RequestCtx)
}

type IServerPool interface {
	AddBackend(IBackend)
	GetBackends() []IBackend
	GetSize() uint32
}

type IBackend interface {
	ProxyHandler(ctx *fasthttp.RequestCtx) error
}
