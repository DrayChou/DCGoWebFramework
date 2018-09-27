package controller

import (
	"fmt"

	"../app"
)

type IndexController struct {
	app.Controller
}

func (p IndexController) Index() {
	fmt.Fprintln(p.Response, p.Request.RequestURI)
	fmt.Fprintln(p.Response, "index.index")
}
