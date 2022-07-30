package controller

import "cncamp/module10/framework"

func SubjectAddController(c *framework.Context) {
	c.SetOkStatus().Json("ok, SubjectAddController")
}

func SubjectListController(c *framework.Context) {
	c.SetOkStatus().Json("ok, SubjectListController")
}

func SubjectDelController(c *framework.Context) {
	c.SetOkStatus().Json("ok, SubjectDelController")
}

func SubjectUpdateController(c *framework.Context) {
	c.SetOkStatus().Json("ok, SubjectUpdateController")
}

func SubjectGetController(c *framework.Context) {
	c.SetOkStatus().Json("ok, SubjectGetController")
}

func SubjectNameController(c *framework.Context) {
	c.SetOkStatus().Json("ok, SubjectNameController")
}
