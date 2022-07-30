package main

import (
	"cncamp/module10/framework"
	"cncamp/module10/framework/middleware"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())

	groupApi := core.Group("simple-web")

	registerRouter(groupApi)
	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Server Start: ", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
}
