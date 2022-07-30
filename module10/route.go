package main

import (
	"cncamp/module10/framework"
	"cncamp/module10/framework/middleware"
)

func registerRouter(groupApi *framework.Group) {

	groupApi.Get("/healthz", HealthzController)
	groupApi.Get("/timeout", TimeoutController)
	groupApi.Get("/user/login", middleware.Cost(), UserController)

	subApi := groupApi.Group("/sub")
	{
		subApi.Use(middleware.Cost())

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
