package middleware

import (
	"cncamp/module10/framework"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Timeout(d time.Duration) framework.HandlerFunc {
	return func(ctx *framework.Context) {

		done := make(chan struct{}, 1)
		panicChan := make(chan any, 1)

		durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			ctx.Next()

			done <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			log.Default().Println(p)
			ctx.SetStatus(http.StatusInternalServerError).Json("panic")
		case <-done:
			fmt.Println("done")
		case <-durationCtx.Done():
			ctx.SetStatus(http.StatusInternalServerError).Json("timeout")
			ctx.SetTimeout()
		}
	}
}
