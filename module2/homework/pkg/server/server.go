package server

import (
	"cncamp/module2/homework/pkg/route"
	"context"
	"net/http"
	"sync"
)

type HttpServer interface {
	route.Router
	Start() error
	Shutdown(ctx context.Context) error
}

type simpleHttpServer struct {
	addr          string
	routerHandler route.RouterHandler
	contextPool   sync.Pool
}

// Route 路由注册
func (shs *simpleHttpServer) Route(method string, path string, handleFunc route.HandleFunc) {
	shs.routerHandler.Route(method, path, handleFunc)
}

func (shs *simpleHttpServer) Start() error {
	http.Handle("/", shs.routerHandler)
	return http.ListenAndServe(shs.addr, nil)
}

func (shs *simpleHttpServer) Shutdown(ctx context.Context) error {
	return nil
}

func NewSimpleHttpServer(addr string) HttpServer {
	return &simpleHttpServer{
		addr:          addr,
		routerHandler: route.NewSimpleRouterHandler(),
	}
}
