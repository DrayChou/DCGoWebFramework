package main

import (
	"fmt"

	"../../DCGoWebFramework"
	"./controller"
	"./lib"
)

func main() {
	lib.Tools()

	fmt.Println("WARNING: This is an example, but not really safe.")

	app := DCGoWebFramework.New()
	app.Set("index", &controller.IndexController{})
	app.Set("site", &controller.SiteController{})
	app.Run(":8888")
}
