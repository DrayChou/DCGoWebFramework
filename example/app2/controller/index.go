package controller

import (
	"../../.."
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type IndexController struct {
	DCGoWebFramework.Controller
}

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

func (p IndexController) Get() {
	mgodb, err1 := p.App.GetDB("main")
	if err1 != nil {
		panic(err1)
	}

	// Collection People
	c := mgodb.(*mgo.Session).DB("test").C("people")
	var results []Person

	// Query All
	err1 = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)
	if err1 != nil {
		panic(err1)
	}

	b, err := json.Marshal(results)
	if err != nil {
		fmt.Fprintln(p.Response, "json.Marshal failed:", err)
		return
	}
	fmt.Fprintln(p.Response, string(b))
}
