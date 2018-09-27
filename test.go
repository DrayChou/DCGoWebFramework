package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type App struct {
	w http.ResponseWriter
}

//控制器
func (this *App) Say() {
	fmt.Fprintf(this.w, "Say called")
}

func test() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/public/") {
			//匹配静态文件服务
			fmt.Fprintf(w, "method %s not found", r.URL.Path)
		} else {
			app := &App{w}
			rValue := reflect.ValueOf(app)
			rType := reflect.TypeOf(app)
			path := strings.Split(r.URL.Path, "/")
			controlName := path[1]
			method, exist := rType.MethodByName(controlName)
			if exist {
				args := []reflect.Value{rValue}
				method.Func.Call(args)
			} else {
				fmt.Fprintf(w, "method %s not found", r.URL.Path)
			}
		}
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
