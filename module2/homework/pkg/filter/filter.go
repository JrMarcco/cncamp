package filter

import (
	"cncamp/module2/homework/pkg/hctx"
	"fmt"
	"time"
)

type HttpFiler func(httpCtx *hctx.HttpContext)

type HttpFilterBuilder func(next HttpFiler) HttpFiler

var _ HttpFilterBuilder = GlobalFilterBuilder

func GlobalFilterBuilder(next HttpFiler) HttpFiler {
	return func(httpCtx *hctx.HttpContext) {
		startTime := time.Now().UnixNano()

		method := httpCtx.Req.Method
		path := httpCtx.Req.URL.Path
		host := httpCtx.Req.Host

		next(httpCtx)

		fmt.Printf("Handle request [%s] [%s] from [%s] cost %d ns\n", method, path, host, time.Now().UnixNano()-startTime)
	}
}
