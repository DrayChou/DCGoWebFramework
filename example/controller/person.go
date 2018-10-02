package controller

import (
	"../../../DCGoWebFramework"
	"bytes"
	"fmt"
	"gopkg.in/mgo.v2"
	"text/template"
	"time"
)

type PersonController struct {
	DCGoWebFramework.Controller
}

func (p PersonController) Index() {
	mgodb, err1 := mgo.Dial("127.0.0.1")
	if err1 != nil {
		panic(err1)
	}
	defer mgodb.Close()

	mgodb.SetMode(mgo.Monotonic, true)
	// Collection People
	c := mgodb.DB("test").C("people")

	// Query All
	var results []Person
	err1 = c.Find(nil).Sort("-timestamp").All(&results)

	if err1 != nil {
		panic(err1)
	}

	items := make(map[interface{}]interface{})
	fmt.Println(results)
	for i := range results {
		fmt.Println(results[i])
		items[i] = results[i]
	}

	data := make(map[interface{}]interface{})
	data["Title"] = "My Index page"
	data["Items"] = items

	fmt.Println(data)

	p.Tpl("person-index", data)
}

func (p PersonController) Add() {
	fmt.Println(p.Request.Method)
	if p.Request.Method != "POST" {
		p.Response.WriteHeader(400)
		return
	}

	name := p.Request.FormValue("Name")
	phone := p.Request.FormValue("Phone")

	mgodb, err1 := mgo.Dial("127.0.0.1")
	if err1 != nil {
		panic(err1)
	}
	defer mgodb.Close()

	c := mgodb.DB("test").C("people")
	err1 = c.Insert(&Person{Name: name, Phone: phone, Timestamp: time.Now()})

	if err1 != nil {
		panic(err1)
	}

	p.Response.Header().Set("Content-Type","text/html")
	var doc bytes.Buffer
	var templateString = "ok, <a href='javascript:history.go(-1);'>return</a>"
	t := template.New("fieldname example")
	t, _ = t.Parse(templateString)
	t.Execute(&doc, nil)
	html := doc.String()
	fmt.Fprint(p.Response, html)

	return
}
