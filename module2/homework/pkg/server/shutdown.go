package server

import (
	"context"
	"fmt"
	"sync"
)

type Hook func(ctx context.Context) error

func BuildShutdownHook(httpServers ...HttpServer) Hook {
	return func(ctx context.Context) error {

		wg := sync.WaitGroup{}
		wg.Add(len(httpServers))

		finished := make(chan struct{})

		for _, httpServer := range httpServers {
			go func(server HttpServer) {
				if err := server.Shutdown(ctx); err != nil {
					fmt.Printf("Fail to shutdown http server: %v \n", err)
				}
				wg.Done()
			}(httpServer)
		}

		go func() {
			wg.Wait()
			// 所有服务正常关闭
			finished <- struct{}{}
		}()

		select {
		case <-ctx.Done():
			// 超时
			fmt.Println("Close servers timeout")
			return fmt.Errorf("Close servers timeout\n")
		case <-finished:
			return nil
		}
	}
}

func WaitForShutdown(hooks ...Hook) {
}
