package internal

import "cncamp/module10/framework"

func HealthzController(ctx *framework.Context) {
	ctx.SetOkStatus().Json("activating")
}
