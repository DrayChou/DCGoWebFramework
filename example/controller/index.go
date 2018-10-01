package controller

import (
	"fmt"

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

	fmt.Println("MySessionID:", DCGoWebFramework.MySessionID)
	DCGoWebFramework.MySessionMgr.SetSessionVal(DCGoWebFramework.MySessionID, "UserInfo", data)

	userinfo, err := DCGoWebFramework.MySessionMgr.GetSessionVal(DCGoWebFramework.MySessionID, "UserInfo")
	fmt.Println("UserInfo:", userinfo, err)

	p.Tpl("index-index", data)
}
