package main

import (
	"cncamp/module10/framework"
	"time"
)

func TimeoutController(ctx *framework.Context) {
	time.Sleep(10 * time.Second)
	ctx.SetOkStatus().Json("ok, TimeoutController")
}
