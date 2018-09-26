package main

import (
	//	"encoding/json"
	"fmt"
	//	"io/ioutil"
	//	"log"
	//	"net/http"

	//	"reflect"
	//	"strings"

	"./app"
	"./controller"
	"./lib"
)

func main() {
	lib.Tools()

	fmt.Println("WARNING: This is an example, but not really safe.")

	application := app.New()
	application.Get("index", &controller.IndexController{})
	application.Get("site", &controller.SiteController{})
	application.Run(":8888")
}
