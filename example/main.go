package main

import (
	"../../DCGoWebFramework"
	app1c "./app1/controller"
	app2c "./app2/controller"
	"github.com/akkuman/parseConfig"
)

func main() {
	app1_config := parseConfig.New("app1/conf/config.json")
	app1 := DCGoWebFramework.New(&app1_config)
	app1.Set("index", &app1c.IndexController{})
	app1.Set("person", &app1c.PersonController{})
	app1.Run()

	app2_config := parseConfig.New("app2/conf/config.json")
	app2 := DCGoWebFramework.New(&app2_config)
	app2.Set("index", &app2c.IndexController{})
	app2.Run()
}
