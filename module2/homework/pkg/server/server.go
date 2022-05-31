package server

import (
	"cncamp/module2/homework/pkg/route"
	"net/http"
)

type HttpServer interface {
	route.Router
	Start() error
}

type simpleHttpServer struct {
	addr          string
	routerHandler route.RouterHandler
}

// Route 路由注册
func (shs *simpleHttpServer) Route(method string, path string, handleFunc route.HandleFunc) {
	shs.routerHandler.Route(method, path, handleFunc)
}

func (shs *simpleHttpServer) Start() error {
	http.Handle("/", shs.routerHandler)
	return http.ListenAndServe(shs.addr, nil)
}

func NewSimpleHttpServer(addr string) HttpServer {
	return &simpleHttpServer{
		addr:          addr,
		routerHandler: route.NewSimpleRouterHandler(),
	}
}
