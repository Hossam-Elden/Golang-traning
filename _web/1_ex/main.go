package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",foo)
	mux.HandleFunc("/bar/",bar)
	mux.HandleFunc("/about",about)
	http.ListenAndServe(":8080", mux)
}

func foo(w http.ResponseWriter , req *http.Request){
	tpl.ExecuteTemplate(w , "foo.gohtml" , req.Method)
}


func bar(w http.ResponseWriter , req *http.Request){
	tpl.ExecuteTemplate(w , "bar.gohtml" , req.URL)
}



func about(w http.ResponseWriter , req *http.Request){
	d:= struct {
		fname string
		lname string
 	}{
		"john",
		"wick",
	}

	tpl.ExecuteTemplate(w , "me.gohtml" , d)
}

