package DCGoWebFramework

import (
	"fmt"
	"github.com/akkuman/parseConfig"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type application struct {
	routes     map[string]interface{}
	app        string
	Config     *parseConfig.Config
	SessionMgr *SessionMgr
}

/**
初始化对象
 */
func New(config *parseConfig.Config) *application {
	// 是否需要初始化 session
	var SessionMgr *SessionMgr
	sessionKey := config.Get("session > key").(string)
	if sessionKey != "" {
		lifeTime := config.Get("session > life_time").(float64)
		if lifeTime == 0 {
			lifeTime = 3600
		}

		SessionMgr = NewSessionMgr(sessionKey, int64(lifeTime))
	}

	app := config.Get("app").(string)
	return &application{
		routes:     make(map[string]interface{}),
		app:        app,
		Config:     config,
		SessionMgr: SessionMgr,
	}
}

func (p *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}

	static_path := p.Config.Get("static_path").(string)
	staticFilePath, err := GetFilePath(upath, static_path)
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
	methodName := strings.Title(strings.ToLower(r.Method))
	pathArr := strings.Split(upath, "/")
	//	fmt.Println("path:", pathArr)
	if len(pathArr) >= 2 && len(pathArr[1]) > 2 {
		controllerName = strings.ToLower(pathArr[1])
	}
	if len(pathArr) >= 3 && len(pathArr[2]) > 2 {
		actionName = strings.Title(pathArr[2])
	}
	fmt.Printf("controller: %s, action: %s, method: %s\n", controllerName, actionName, methodName)

	route, ok := p.routes[controllerName]
	if !ok {
		http.Error(w, "Controller Not Found", http.StatusNotFound)
		return
	}

	var action string
	_, exist := reflect.TypeOf(route).MethodByName(actionName)
	if exist {
		action = actionName
	} else {
		_, exist = reflect.TypeOf(route).MethodByName(methodName)
		if exist {
			action = methodName
		}
	}

	fmt.Printf("action: %s\n", action)
	if exist {
		ele := reflect.ValueOf(route).Elem()
		ele.FieldByName("Request").Set(reflect.ValueOf(r))
		ele.FieldByName("Response").Set(reflect.ValueOf(w))
		ele.FieldByName("App").Set(reflect.ValueOf(p))
		ele.MethodByName(action).Call([]reflect.Value{})
	} else {
		fmt.Fprintf(w, "methodName %s not found", upath)
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

func (p *application) GetDB(db string) (interface{}, error) {
	var db_conf map[string]interface{}
	db_conf = p.Config.Get("database > " + db).(map[string]interface{})
	fmt.Println("db_conf > ", db, db_conf)

	db_type := db_conf["type"].(string)
	db_host := db_conf["host"].(string)
	db_port := db_conf["port"].(string)
	db_usernamet := db_conf["username"].(string)
	db_password := db_conf["password"].(string)

	switch db_type {
	case "mongodb":
		{
			// mongodb://myuser:mypass@localhost:40001/mydb
			var db_url string
			if (len(db_usernamet) > 1) && (len(db_password) > 1) {
				db_url += db_usernamet + ":" + db_password + "@"
			}
			db_url += db_host + ":" + db_port
			mgodb, err1 := mgo.Dial(db_url)
			if err1 != nil {
				fmt.Println("mgo err1", db_url, mgodb, err1)
				panic(err1)
			}

			//defer mgodb.Close()

			mgodb.SetMode(mgo.Monotonic, true)

			return mgodb, nil
		}
	}

	return nil, mgo.ErrNotFound
}

func (p *application) Run() error {
	p.printRoutes()

	host := p.Config.Get("host").(string)
	port := p.Config.Get("port").(string)
	addr := host + ":" + port
	fmt.Printf("listen on %s\n", addr)

	return http.ListenAndServe(addr, p)
}
