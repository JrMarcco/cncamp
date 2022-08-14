package internal

import (
	"cncamp/module10/framework"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func prometheusHandler() framework.HandlerFunc {
	prometheusHandler := promhttp.Handler()
	return func(ctx *framework.Context) {
		prometheusHandler.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
	}
}

func RegisterRouter(core *framework.Core) {

	core.Get("/metrics", prometheusHandler())

	core.Get("/healthz", HealthzController)
	core.Get("/time", TimeController)
	core.Get("/user/login", UserController)

	subApi := core.Group("/sub")
	{
		subApi.Get("/:id", SubjectGetController)
		subApi.Put("/:id", SubjectUpdateController)
		subApi.Delete("/:id", SubjectDelController)
		subApi.Get("/list/all", SubjectListController)

		subInnerApi := subApi.Group("/info")
		{
			subInnerApi.Get("/name", SubjectNameController)
		}
	}
}
