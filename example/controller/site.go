package controller

import (
	"fmt"

	"../../../DCGoWebFramework"
)

type SiteController struct {
	DCGoWebFramework.Controller
}

func (p SiteController) Index() {
	fmt.Fprintln(p.Response, p.Request.RequestURI)
	fmt.Fprintln(p.Response, "site.index")
}
