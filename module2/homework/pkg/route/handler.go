package route

import (
	"cncamp/module2/homework/pkg/hctx"
	"fmt"
	"sync"
)

type RouterHandler interface {
	HandleHttp(httpCtx *hctx.HttpContext)
	Router
}

type SimpleRouterHandler struct {
	handlers sync.Map
}

func (srh *SimpleRouterHandler) HandleHttp(httpCtx *hctx.HttpContext) {
	handlerKey := fmt.Sprintf("%s %s", httpCtx.Req.Method, httpCtx.Req.URL.Path)

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
