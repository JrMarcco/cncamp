package route

import (
	"cncamp/module2/homework/pkg/hctx"
)

type HandleFunc func(httpCtx *hctx.HttpContext)

type Router interface {
	Route(method string, path string, handleFunc HandleFunc)
}
