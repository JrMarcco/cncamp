package route

import (
	"cncamp/module2/homework/pkg/hctx"
	"fmt"
	"net/http"
	"sync"
)

type RouterHandler interface {
	http.Handler
	Router
}

type SimpleRouterHandler struct {
	handlers sync.Map
}

func (srh *SimpleRouterHandler) ServeHTTP(rspWriter http.ResponseWriter, req *http.Request) {
	handlerKey := fmt.Sprintf("%s %s", req.Method, req.URL.Path)

	httpCtx := hctx.NewHttpContext(rspWriter, req)
	if handler, ok := srh.handlers.Load(handlerKey); ok {
		handler.(HandleFunc)(httpCtx)
		return
	}
	httpCtx.NotFound()
}

func (srh *SimpleRouterHandler) Route(method string, path string, handleFunc HandleFunc) {
	handlerKey := fmt.Sprintf("%s %s", method, path)
	if _, ok := srh.handlers.Load(handlerKey); ok {
		// 不允许重复注册路由
		panic(fmt.Sprintf("Repeat registered for path: %s\n", handlerKey))
	}
	srh.handlers.Store(handlerKey, handleFunc)
}

func NewSimpleRouterHandler() *SimpleRouterHandler {
	return &SimpleRouterHandler{}
}
