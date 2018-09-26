package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"./docgowebframework"
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func filename(name string) string {
	fmt.Println("name: " + name)

	if name == "" {
		name = "query"
	}

	// A little security, but don't be fooled, this still isn't really safe.
	name = strings.Replace(name, ".", "", -1)
	name = strings.TrimPrefix(name, "/")
	name += ".txt"
	fmt.Println("filename: " + name)
	return name
}

func get(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Query: ", req.URL.Query())
	fmt.Println("Path: ", req.URL.Path)
	fmt.Println("RawPath: ", req.URL.RawPath)

	fmt.Println("get2: ", req.URL.Query().Get("name"))
	fmt.Println("get3: ", req.Form)
	fmt.Println("get4: ", req.Form.Get("name"))

	// Example of fetching specific Query Param.
	name := filename(req.Form.Get("name"))
	body, err := ioutil.ReadFile(name)

	if err == nil {
		fmt.Fprintf(res, string(body[:]))
	} else {
		fmt.Fprintf(res, "Error: ", err)
	}
}

func post(res http.ResponseWriter, req *http.Request) {
	// Example of fetching specific Query Param.
	name := filename(req.Form.Get("name"))

	// Example of creating JSON string.
	body, err := json.Marshal(req.Form)
	if err != nil {
		fmt.Fprintf(res, "ERROR: ", err)
	} else {
		ioutil.WriteFile(name, body, 0600)
		fmt.Fprintf(res, string(body[:]))
	}
}

func bad(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "GO AWAY", req.Method, req.URL.Path)
}

func handler(res http.ResponseWriter, req *http.Request) {
	// Example of parsing GET or POST Query Params.
	req.ParseForm()

	// Example of handling POST request.
	switch req.Method {
	case "POST":
		post(res, req)
	// Example of handling GET request.
	case "GET":
		get(res, req)
	default:
		bad(res, req)
	}
}

func main() {
	docgowebframework.test()

	fmt.Println("WARNING: This is an example, but not really safe.")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", Log(http.DefaultServeMux))
}
