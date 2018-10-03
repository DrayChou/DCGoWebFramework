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
	App       *application
	SessionID string
}

func (p *Controller) SessionStart() {
	p.SessionID = p.App.SessionMgr.CheckCookieValid(p.Response, p.Request)
	fmt.Println("CheckCookieValid:", p.SessionID)

	if len(p.SessionID) < 1 {
		//fmt.Println("StartSession:")
		p.SessionID = p.App.SessionMgr.StartSession(p.Response, p.Request)
	}
	fmt.Println("MySessionID:", p.SessionID)
}

func (p *Controller) SetSessionVal(key interface{}, value interface{}) {
	p.App.SessionMgr.SetSessionVal(p.SessionID, key, value)
}

func (p *Controller) GetSessionVal(key interface{}) (interface{}, bool) {
	return p.App.SessionMgr.GetSessionVal(p.SessionID, key)
}

func (p Controller) Tpl(path string, data map[interface{}]interface{}) {
	basePath, _ := os.Getwd()

	template_path := p.App.Config.Get("template_path").(string)
	tempPath := filepath.Join(basePath, p.App.app, template_path, path+".html")
	fmt.Println("tempPath:", tempPath)
	if !IsPathExist(tempPath) {
		return
	}

	t, err := template.ParseFiles(tempPath)
	if err != nil {
		fmt.Println("ParseFiles", err)
	}

	err = t.Execute(p.Response, data)
	if err != nil {
		fmt.Println("Execute", err)
	}
}
