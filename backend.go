package main

import (
	"github.com/daniil-oliynyk/gobalancer/config"
	"github.com/daniil-oliynyk/gobalancer/utils"
	"github.com/valyala/fasthttp"
)

type IBackend interface {
	ProxyHandler(ctx *fasthttp.RequestCtx) error
}

type Backend struct {
	proxy *fasthttp.HostClient
	addr  string
}

func (b *Backend) ProxyHandler(ctx *fasthttp.RequestCtx) error {

	logger := utils.GetLogger()

	req := &ctx.Request
	res := &ctx.Response
	clientIP := ctx.RemoteIP()
	b.PrepareRequest(req, clientIP)

	err := b.proxy.Do(req, res)
	if err != nil {
		logger.Error().Err(err).Msg("Error when proxying request")

		res.SetStatusCode(fasthttp.StatusInternalServerError)
		res.SetConnectionClose()
		res.Header.Set("Content-Type", "application/json")
		res.SetBody([]byte(`"message":"` + err.Error() + `"}`))
		return err
	}

	return nil
}
func (b *Backend) PrepareRequest(req *fasthttp.Request, clientIP []byte) {
	req.URI().SetScheme("http")
	req.URI().SetHost(b.addr)
	req.Header.SetBytesKV([]byte("X-Forwarded-For"), clientIP)
}

func NewBackend(backend config.Backend) IBackend {
	// More fields can be added later in config such as MaxConns
	proxyHostClient := &fasthttp.HostClient{
		Addr: backend.Url,
	}

	return &Backend{
		proxy: proxyHostClient,
		addr:  backend.Url,
	}

}
