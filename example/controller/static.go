package controller

import (
	"fmt"

	"../../../DCGoWebFramework"
)

type StaticController struct {
	DCGoWebFramework.Controller
}

func (p StaticController) Index() {
	fmt.Fprintln(p.Response, p.Request.RequestURI)
	fmt.Fprintln(p.Response, "Static.index")
}
