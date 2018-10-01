package DCGoWebFramework

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Controller struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func (p Controller) Tpl(path string, data map[interface{}]interface{}) {
	basePath, _ := os.Getwd()

	tempPath := filepath.Join(basePath, "template", path+".html")
	fmt.Println("tempPath:", tempPath)
	if !IsPathExist(tempPath) {
		return
	}

	fmt.Println("IsPathExist:", 1)

	t, err := template.ParseFiles(tempPath)
	if err != nil {
		fmt.Println("ParseFiles", err)
	}

	err = t.Execute(p.Response, data)
	if err != nil {
		fmt.Println("Execute", err)
	}
}
