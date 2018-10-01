package DCGoWebFramework

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var MySessionKey = "DCGoWebFramework-sid"
var MySessionMgr *SessionMgr
var MySessionID string

type Controller struct {
	Response   http.ResponseWriter
	Request    *http.Request
	SessionMgr *SessionMgr
	SessionID  string
}

func (p Controller) SessionStart() {
	c1, err := p.Request.Cookie(MySessionKey)
	fmt.Println("Cookie:", c1, err)

	MySessionMgr = NewSessionMgr(MySessionKey, 3600)
	MySessionID = MySessionMgr.CheckCookieValid(p.Response, p.Request)
	fmt.Println("CheckCookieValid:", MySessionID)

	if len(MySessionID) < 1 {
		fmt.Println("StartSession:")
		MySessionID = MySessionMgr.StartSession(p.Response, p.Request)
	}
	fmt.Println("MySessionID:", MySessionID)

	// 这个值带不到外面去，废弃
	p.SessionMgr = MySessionMgr
	p.SessionID = MySessionID

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
