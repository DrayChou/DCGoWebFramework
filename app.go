package DCGoWebFramework

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
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}

	controllerName := "index"
	actionName := "Index"
	path := strings.Split(upath, "/")
	fmt.Println("path:", path)
	if len(path) >= 2 {
		controllerName = strings.ToLower(path[1])
	}
	if len(path) >= 3 {
		actionName = strings.Title(path[2])
	}
	fmt.Printf("%s %s\n", controllerName, actionName)

	if controllerName == "favicon.ico" {
		//匹配静态文件服务
		fmt.Println("static 001", upath)
		fmt.Println("static 002", http.Dir("../static/favicon.ico"))
		fmt.Println("static 003", http.FileServer(http.Dir("../static/favicon.ico")))

		//		http.Handle("/static/favicon.ico", http.StripPrefix("/static/favicon.ico", http.FileServer(http.Dir("../static/favicon.ico"))))
		http.StripPrefix("/static/favicon.ico", http.FileServer(http.Dir("../static/favicon.ico")))
		return
	}

	if controllerName == "static" {
		//匹配静态文件服务
		fmt.Println("static %s was found", upath)
		fmt.Println("static 001", upath)
		fmt.Println("static 002", http.Dir("../static/favicon.ico"))
		fmt.Println("static 003", http.FileServer(http.Dir("../static/favicon.ico")))

		//		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
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
		fmt.Fprintf(w, "method %s not found", upath)
	}
}

func (p *application) StaticServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StaticServer 001", r.URL.Path)
	http.StripPrefix("/static", http.FileServer(http.Dir("../static/"))).ServeHTTP(w, r)
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

	return http.ListenAndServe(addr, p)
}
