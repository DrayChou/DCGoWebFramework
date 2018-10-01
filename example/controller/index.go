package controller

import (
	"../../../DCGoWebFramework"
)

type IndexController struct {
	DCGoWebFramework.Controller
}

func (p IndexController) Index() {

	data := make(map[interface{}]interface{})
	data["Title"] = "My Index page"
	data["Items"] = []string{
		"My photos",
		"My blog",
	}
	p.SessionStart()
	p.SessionMgr.SetSessionVal(p.SessionID, "UserInfo", data)

	p.Tpl("index-index", data)
}
