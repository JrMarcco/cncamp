package internal

import (
	"cncamp/module10/framework"
	"cncamp/module10/framework/middleware"
	"cncamp/module10/internal/controller"
)

func RegisterRouter(groupApi *framework.Group) {

	groupApi.Get("/healthz", controller.HealthzController)
	groupApi.Get("/timeout", controller.TimeoutController)
	groupApi.Get("/user/login", middleware.Cost(), controller.UserController)

	subApi := groupApi.Group("/sub")
	{
		subApi.Use(middleware.Cost())

		subApi.Get("/:id", controller.SubjectGetController)
		subApi.Put("/:id", controller.SubjectUpdateController)
		subApi.Delete("/:id", controller.SubjectDelController)
		subApi.Get("/list/all", controller.SubjectListController)

		subInnerApi := subApi.Group("/info")
		{
			subInnerApi.Get("/name", controller.SubjectNameController)
		}
	}
}
