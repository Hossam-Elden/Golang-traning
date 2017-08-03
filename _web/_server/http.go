package main

import (
	"net/http"
	"io"
	"log"
	"html/template"
)
// normal static route sending simple data to it
func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo run")
}

// route for file text and send data from the new function

func new(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("new.hos")
	if err != nil {
		log.Fatalln("error parsing ", err)
	}

	err = tpl.ExecuteTemplate(res, "new.hos", "hellllo from the other side")
	if err != nil {
		log.Fatalln("error execute ", err)
	}

}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/new/", new)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("server error", err)
	}
}
