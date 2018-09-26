package controller

import (
	"fmt"

	"../app"
)

type IndexController struct {
	app.Controller
}

func (p IndexController) Index() {
	fmt.Fprint(p.Response, p.Request.RequestURI)
}
