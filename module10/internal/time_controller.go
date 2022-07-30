package internal

import (
	"cncamp/module10/framework"
	"math/rand"
	"time"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func TimeController(ctx *framework.Context) {
	executionTimer := NewExecutionTimer()
	defer executionTimer.ObserveTotal()

	time.Sleep(time.Millisecond * time.Duration(randInt(10, 2000)))

	ctx.SetOkStatus().Json("ok, TimeController")
}
