package controller

import (
	"fmt"

	"../../../DCGoWebFramework"
)

type IndexController struct {
	DCGoWebFramework.Controller
}

func (p IndexController) Index() {
	fmt.Fprintln(p.Response, p.Request.RequestURI)
	fmt.Fprintln(p.Response, "index.index")
}
