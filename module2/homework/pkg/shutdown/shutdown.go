package shutdown

import (
	"cncamp/module2/homework/pkg/server"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Hook func(ctx context.Context) error

func BuildShutdownHook(httpServers ...server.HttpServer) Hook {
	return func(ctx context.Context) error {

		wg := sync.WaitGroup{}
		wg.Add(len(httpServers))

		finished := make(chan struct{})

		for _, httpServer := range httpServers {
			go func(server server.HttpServer) {
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
			fmt.Println("Close servers timeout")
			return fmt.Errorf("Close servers timeout\n")
		case <-finished:
			return nil
		}
	}
}

func Wait(hooks ...Hook) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, ShutdownSignals...)

	select {
	case <-signalChannel:
		time.AfterFunc(5*time.Minute, func() {
			fmt.Println("Shutdown timeout, force exit now")
			os.Exit(1)
		})

		for _, hook := range hooks {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			if err := hook(ctx); err != nil {
				fmt.Printf("Fail to run shutdown hook: %v\n", err)
			}
			cancel()
		}
	}
}
