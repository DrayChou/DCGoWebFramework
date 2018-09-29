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

	marrage_info := map[string]string{
		"HanMeimei": "1",
		"LiLei":     "0",
	}
	DCGoWebFramework.Controller.Temp("index.index", *marrage_info)

}
