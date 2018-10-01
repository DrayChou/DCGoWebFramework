package controller

import (
	"../../../DCGoWebFramework"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PersonController struct {
	DCGoWebFramework.Controller
}

func (p PersonController) Index() {
	session, err1 := mgo.Dial("127.0.0.1")
	if err1 != nil {
		panic(err1)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	// Collection People
	c := session.DB("test").C("people")

	// Query All
	var results []Person
	err1 = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err1 != nil {
		panic(err1)
	}

	//for i := range results {
	//	fmt.Println(results[i])
	//	var byteSlice []res = *(*[]byte)(unsafe.Pointer(sb))
	//}
	//
	//fmt.Fprintf(p.Response, res)

	p.Tpl("site-index", results)
}
