package internal

import "cncamp/module10/framework"

func UserController(ctx *framework.Context) {
	ctx.SetOkStatus().Json("ok, UserController")
}
