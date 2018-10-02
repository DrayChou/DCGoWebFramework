package DCGoWebFramework

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

var MySessionKey = "DCGoWebFramework-sid"
var MySessionMgr *SessionMgr

type application struct {
	routes map[string]interface{}
}

/**
初始化对象
 */
func New(sessionKey string) *application {
	if len(sessionKey) > 1 {
		MySessionKey = sessionKey
	}

	MySessionMgr = NewSessionMgr(MySessionKey, 3600)

	//fmt.Println("New", MySessionKey, MySessionMgr, MySessionID)

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

	staticFilePath, err := GetFilePath(upath)
	if !err {
		//匹配静态文件服务
		body, err := ioutil.ReadFile(staticFilePath)
		if err == nil {
			fmt.Fprintf(w, string(body[:]))
		} else {
			fmt.Fprintf(w, "Error: ", err)
		}
		return
	}

	controllerName := "index"
	actionName := "Index"
	pathArr := strings.Split(upath, "/")
	//	fmt.Println("path:", pathArr)
	if len(pathArr) >= 2 && len(pathArr[1]) > 2 {
		controllerName = strings.ToLower(pathArr[1])
	}
	if len(pathArr) >= 3 && len(pathArr[2]) > 2 {
		actionName = strings.Title(pathArr[2])
	}
	fmt.Printf("controllerName: %s ,actionName: %s\n", controllerName, actionName)

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
		//ele.MethodByName("SessionStart").Call([]reflect.Value{})
		ele.MethodByName(actionName).Call([]reflect.Value{})
	} else {
		fmt.Fprintf(w, "method %s not found", upath)
	}
}

func (p *application) printRoutes() {
	for route, controller := range p.routes {
		ele := reflect.ValueOf(controller).Type().String()
		fmt.Printf("%s %s\n", route, ele)
	}
}

func (p *application) Set(route string, controller interface{}) {
	p.routes[route] = controller
}

func (p *application) Run(addr string) error {
	p.printRoutes()
	fmt.Printf("listen on %s\n", addr)

	return http.ListenAndServe(addr, p)
}
