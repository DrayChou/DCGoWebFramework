package controller

import (
	"../../.."
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

func (p IndexController) GET() {
	p.SessionStart()

	data := make(map[interface{}]interface{})
	data["Title"] = "My Index page"
	data["Items"] = []string{
		"My photos",
		"My blog",
	}

	fmt.Println("MySessionID:", p.SessionID)
	p.SetSessionVal("UserInfo", data)

	userinfo, err := p.GetSessionVal("UserInfo")
	fmt.Println("UserInfo:", userinfo, err)

	mgodb, err1 := p.App.GetDB("main")
	if err1 != nil {
		panic(err1)
	}

	// Collection People
	c := mgodb.(*mgo.Session).DB("test").C("people")

	//// Index
	//index := mgo.Index{
	//	Key:        []string{"name", "phone"},
	//	Unique:     false,
	//	DropDups:   true,
	//	Background: true,
	//	Sparse:     true,
	//}
	//
	//err1 = c.EnsureIndex(index)
	//if err1 != nil {
	//	panic(err1)
	//}

	// Insert Datas
	err1 = c.Insert(
		&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
		&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})

	if err1 != nil {
		panic(err1)
	}

	// Query One
	result := Person{}
	err1 = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("Phone", result)

	// Query All
	var results []Person
	err1 = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err1 != nil {
		panic(err1)
	}
	fmt.Println("Results All: ", results)

	// Update
	colQuerier := bson.M{"name": "Ale"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	err1 = c.Update(colQuerier, change)
	if err1 != nil {
		panic(err1)
	}

	// Query All
	err1 = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err1 != nil {
		panic(err1)
	}
	fmt.Println("Results All: ", results)

	p.Tpl("index-index", data)
}
