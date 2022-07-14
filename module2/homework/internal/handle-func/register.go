package handle_func

import (
	"cncamp/module2/homework/pkg/hctx"
	"fmt"
	"os"
	"strings"
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
	rawQuery := httpCtx.Req.URL.RawQuery
	if rawQuery != "" {
		key = strings.Split(rawQuery, "=")[0]
	}
	if envVar, ok := os.LookupEnv(key); ok {
		httpCtx.SetRspHeader(key, envVar)
		httpCtx.WriteOkStr(envVar)
		return
	}
	httpCtx.WriteOkStr("Env var not found")
}

var HealthzHF = func(httpCtx *hctx.HttpContext) {
	httpCtx.WriteOkStr(fmt.Sprintln("activating"))
}
