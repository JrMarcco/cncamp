package main

import (
	handleFunc "cncamp/module2/homework/internal/handle-func"
	"cncamp/module2/homework/pkg/filter"
	"cncamp/module2/homework/pkg/server"
	"cncamp/module2/homework/pkg/shutdown"
	"net/http"
)

func main() {
	httpServer := server.NewSimpleHttpServer(":8080", filter.GlobalFilterBuilder)

	httpServer.Route(http.MethodGet, "/header", handleFunc.HeaderHF)
	httpServer.Route(http.MethodGet, "/envVar", handleFunc.EnvVarHF)
	httpServer.Route(http.MethodGet, "/healthz", handleFunc.HealthzHF)

	go func() {
		if err := httpServer.Start(); err != nil {
			panic(err)
		}
	}()

	shutdown.Wait(
		// 关闭服务前的操作，例如：摘除流量、拒绝请求等
		shutdown.BuildShutdownHook(httpServer),
		// 关闭服务后的操作，例如：资源关闭等
	)
}
