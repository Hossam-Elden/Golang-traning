package main

import (
	"net/http"
	"html/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templ/form.html")
	t.Execute(w, r.FormValue("comment"))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templ/xss_form.html")
	t.Execute(w, nil)
}

func main() {
	 server := http.Server{
		 Addr:":8080",
	 }
	http.HandleFunc("/form", process)
	http.HandleFunc("/xss_in" , form)

	server.ListenAndServe()
}

