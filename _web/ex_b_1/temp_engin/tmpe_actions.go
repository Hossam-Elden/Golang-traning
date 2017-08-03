package main

import (
	"net/http"
	"html/template"
	"math/rand"
	"time"
)

func process(w http.ResponseWriter, r *http.Request){
	t:= template.Must(template.ParseFiles("tmpl.html"))
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}



func ranger(w http.ResponseWriter, r *http.Request){
	ran:=template.Must(template.ParseFiles("templ/range.html"))
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	ran.Execute(w, daysOfWeek)
}



func main() {
	server:=http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/action" , process)
	http.HandleFunc("/range", ranger)
	server.ListenAndServe()
}