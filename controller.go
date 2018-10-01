package DCGoWebFramework

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Controller struct {
	Response  http.ResponseWriter
	Request   *http.Request
	SessionID string
}

func (p *Controller) SessionStart() {
	p.SessionID = MySessionMgr.CheckCookieValid(p.Response, p.Request)
	fmt.Println("CheckCookieValid:", p.SessionID)

	if len(p.SessionID) < 1 {
		//fmt.Println("StartSession:")
		p.SessionID = MySessionMgr.StartSession(p.Response, p.Request)
	}
	fmt.Println("MySessionID:", p.SessionID)
}

func (p *Controller) SetSessionVal(key interface{}, value interface{}) {
	MySessionMgr.SetSessionVal(p.SessionID, key, value)
}

func (p *Controller) GetSessionVal(key interface{}) (interface{}, bool) {
	return MySessionMgr.GetSessionVal(p.SessionID, key)
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
