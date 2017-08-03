package main

import (
	"net/http"
	"html/template"
)

func process (w http.ResponseWriter, r *http.Request)  {
	t:= template.Must(template.ParseFiles("tmpl.html"))
	t.Execute(w, "hello world")

}

func main() {
	server:=http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/tmpl" ,process)
	server.ListenAndServe()
}
