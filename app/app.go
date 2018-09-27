package app

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type application struct {
	routes map[string]interface{}
}

func New() *application {
	return &application{
		routes: make(map[string]interface{}),
	}
}

func (p *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	controllerName := "index"
	actionName := "Index"
	path := strings.Split(r.URL.Path, "/")
	fmt.Println("path:", path)
	if len(path) >= 2 {
		controllerName = strings.ToLower(path[1])
	}
	if len(path) >= 3 {
		actionName = strings.Title(path[2])
	}
	fmt.Printf("%s %s\n", controllerName, actionName)

	if controllerName == "static" {
		//匹配静态文件服务
		fmt.Println("static %s was found", r.URL.Path)
		http.FileServer(http.Dir("../static"))
		return
	}

	route, ok := p.routes[controllerName]
	if !ok {
		http.Error(w, "Controller Not Found", http.StatusNotFound)
		return
	}

	_, exist := reflect.TypeOf(route).MethodByName(actionName)
	if exist {
		ele := reflect.ValueOf(route).Elem()
		ele.FieldByName("Request").Set(reflect.ValueOf(r))
		ele.FieldByName("Response").Set(reflect.ValueOf(w))
		ele.MethodByName(actionName).Call([]reflect.Value{})
	} else {
		fmt.Fprintf(w, "method %s not found", r.URL.Path)
	}
}

func (p *application) printRoutes() {
	for route, controller := range p.routes {
		ele := reflect.ValueOf(controller).Type().String()
		fmt.Printf("%s %s\n", route, ele)
	}
}

func (p *application) Get(route string, controller interface{}) {
	p.routes[route] = controller
}

func (p *application) Run(addr string) error {
	p.printRoutes()
	fmt.Printf("listen on %s\n", addr)

	http.Handle("/static/", http.FileServer(http.Dir("../static")))

	return http.ListenAndServe(addr, p)
}
