package main

import (
	handle_func "cncamp/module2/homework/internal/handle-func"
	"cncamp/module2/homework/pkg/server"
	"log"
)

func main() {
	httpServer := server.NewSimpleHttpServer(":8080")

	httpServer.Route("GET", "/header", handle_func.HeaderHF)
	httpServer.Route("GET", "/envVar", handle_func.EnvVarHF)
	httpServer.Route("GET", "/healthz", handle_func.HealthzHF)

	if err := httpServer.Start(); err != nil {
		log.Fatalln(err)
	}
}
