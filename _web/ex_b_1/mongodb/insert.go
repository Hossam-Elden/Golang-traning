package main

import (
	"net/http"
	"fmt"
	"html/template"
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type Wish struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Description string
}

func form(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("insert.html"))
	tmpl.Execute(w, nil)
}
func insert(w http.ResponseWriter, r *http.Request) {
	session, err :=mgo.Dial("localhost")
	if err!=nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic,true)
	d :=session.DB("test").C("mytest")



	r.ParseForm()
	doc := Wish{
		bson.NewObjectId(),
		r.FormValue("name"),
		 r.FormValue("description"),
	}

	err = d.Insert(&doc)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, doc)
}

func main() {

	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/form", form)
	http.HandleFunc("/insert", insert)
	server.ListenAndServe()

}
