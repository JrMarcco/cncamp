package handle_func

import (
	"cncamp/module2/homework/pkg/hctx"
	"os"
)

// HeaderHF 将 request 中带的 header 写入 response header
var HeaderHF = func(httpCtx *hctx.HttpContext) {
	for key, value := range httpCtx.Req.Header {
		if len(value) > 0 {
			httpCtx.SetRspHeader(key, value[0])
		}
	}
}

// EnvVarHF 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
var EnvVarHF = func(httpCtx *hctx.HttpContext) {
	key := "VERSION"
	httpCtx.SetRspHeader(key, os.Getenv(key))
}

var HealthzHF = func(httpCtx *hctx.HttpContext) {
	httpCtx.WriteOkStr("activating")
}
