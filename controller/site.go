package controller

import (
	"fmt"

	"../app"
)

type SiteController struct {
	app.Controller
}

func (p SiteController) Index() {
	fmt.Fprintln(p.Response, p.Request.RequestURI)
	fmt.Fprintln(p.Response, "site.index")
}
