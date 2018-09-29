package DCGoWebFramework

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"./utils"
)

type Controller struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func (p Controller) Temp(path string, data map[string]string) {
	basePath, _ := os.Getwd()
	tempPath := basePath + "/template/" + path + ".template"

	fmt.Println("tempPath:", tempPath)
	if utils.IsPathExist(tempPath) {
		fmt.Println("IsPathExist:", 1)
		t := template.New(path)
		//		template.Must(t.ParseFiles(tempPath))
		t.ExecuteTemplate(p.Response, tempPath, data)
	}
}
