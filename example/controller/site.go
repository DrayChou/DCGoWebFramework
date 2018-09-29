package controller

import (
	"fmt"

	"../../../DCGoWebFramework"
)

type SiteController struct {
	DCGoWebFramework.Controller
}

func (p SiteController) Index() {
	data := make(map[interface{}]interface{})
	data["Title"] = "My Site page"
	data["Items"] = []string{
		"My photos",
		"My blog",
	}

	p.Tpl("site-index", data)
}
