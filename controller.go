package DCGoWebFramework

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Controller struct {
	Response   http.ResponseWriter
	Request    *http.Request
	SessionMgr *SessionMgr
	SessionID  string
}

func (p Controller) init() {
	p.SessionMgr = NewSessionMgr("TestCookieName", 3600)
	p.SessionID = p.SessionMgr.StartSession(p.Response, p.Request)
	fmt.Println("SessionID:", p.SessionID)
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
