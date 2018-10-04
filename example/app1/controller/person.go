package controller

import (
	"../../.."
	"bytes"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"text/template"
	"time"
)

type PersonController struct {
	DCGoWebFramework.Controller
}

func (p PersonController) Get() {
	data := make(map[interface{}]interface{})
	data["Title"] = "My Index page"

	mgodb, err1 := p.App.GetDB("main")
	if err1 != nil {
		panic(err1)
	}

	// Collection People
	c := mgodb.(*mgo.Session).DB("test").C("people")

	items := make(map[interface{}]interface{})

	var results []Person
	fmt.Println(p.Request.URL)
	pathArr := strings.Split(p.Request.RequestURI, "/")
	if len(pathArr) == 3 {
		id := string(pathArr[2])
		// db.getCollection('people').find({"_id":ObjectId("5bb250d978422140d307809c")})
		if strings.ToLower(id) != "index" {
			fmt.Println(id)

			// Query All
			err1 = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Sort("-timestamp").All(&results)
			if err1 != nil {
				panic(err1)
			}

			fmt.Println(results)
			items[0] = results[0]

			data["Items"] = items
			p.Tpl("person-index", data)
			return
		}
	}

	// Query All
	err1 = c.Find(nil).Sort("-timestamp").All(&results)
	if err1 != nil {
		panic(err1)
	}

	//fmt.Println(results)
	for i := range results {
		//fmt.Println(results[i])
		items[i] = results[i]
	}
	data["Items"] = items
	p.Tpl("person-index", data)
}

func (p PersonController) Post() {
	fmt.Println(p.Request.Method)
	if p.Request.Method != "POST" {
		p.Response.WriteHeader(400)
		return
	}

	name := p.Request.FormValue("Name")
	phone := p.Request.FormValue("Phone")

	mgodb, err1 := p.App.GetDB("main")
	if err1 != nil {
		panic(err1)
	}

	// Collection People
	c := mgodb.(*mgo.Session).DB("test").C("people")
	err1 = c.Insert(&Person{Name: name, Phone: phone, Timestamp: time.Now()})

	if err1 != nil {
		panic(err1)
	}

	p.Response.Header().Set("Content-Type", "text/html")
	var doc bytes.Buffer
	var templateString = "ok, <a href='javascript:history.go(-1);'>return</a>"
	t := template.New("fieldname example")
	t, _ = t.Parse(templateString)
	t.Execute(&doc, nil)
	html := doc.String()
	fmt.Fprint(p.Response, html)

	return
}
