package server

import (
	"cncamp/module2/homework/pkg/filter"
	"cncamp/module2/homework/pkg/hctx"
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
	addr            string
	routerHandler   route.RouterHandler
	rootHttpFilter  filter.HttpFiler
	httpContextPool sync.Pool
}

func (shs *simpleHttpServer) ServeHTTP(rspWriter http.ResponseWriter, req *http.Request) {
	httpCtx := shs.httpContextPool.Get().(*hctx.HttpContext)
	defer func() {
		shs.httpContextPool.Put(httpCtx)
	}()

	httpCtx.Reset(rspWriter, req)
	shs.rootHttpFilter(httpCtx)
}

// Route 路由注册
func (shs *simpleHttpServer) Route(method string, path string, handleFunc route.HandleFunc) {
	shs.routerHandler.Route(method, path, handleFunc)
}

func (shs *simpleHttpServer) Start() error {
	return http.ListenAndServe(shs.addr, shs)
}

func (shs *simpleHttpServer) Shutdown(ctx context.Context) error {
	// 关闭逻辑
	return nil
}

func NewSimpleHttpServer(addr string, httpFilterBuilders ...filter.HttpFilterBuilder) HttpServer {
	simpleRouteHandler := route.NewSimpleRouterHandler()

	var root filter.HttpFiler = simpleRouteHandler.HandleHttp
	for i := len(httpFilterBuilders) - 1; i >= 0; i-- {
		httpFilterBuilder := httpFilterBuilders[i]
		root = httpFilterBuilder(root)
	}

	return &simpleHttpServer{
		addr:           addr,
		routerHandler:  simpleRouteHandler,
		rootHttpFilter: root,
		httpContextPool: sync.Pool{
			New: func() any {
				return hctx.New()
			},
		},
	}
}
