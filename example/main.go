package main

import (
	"../../DCGoWebFramework"
	"./controller"
)

var MySessionKey = "DCGoWebFramework-sid"

func main() {
	app := DCGoWebFramework.New(MySessionKey)
	app.Set("index", &controller.IndexController{})
	app.Set("person", &controller.PersonController{})
	app.Run(":8888")
}
