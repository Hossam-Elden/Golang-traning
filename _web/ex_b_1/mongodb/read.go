package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"net/http"
	"html/template"
	"log"
)

type Task struct {
	Description string
	Due         time.Time
}

type Category struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Description string
	Tasks       []Task
}

type poll struct {
	Options []string
}
func reading(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	C := session.DB("test").C("mytest")

	iter := C.Find(nil).Iter()
	var  p  Category
	// store the data in slice to send it to html templ
	d := []string{}

	for iter.Next(&p) {
		d = append(d, p.Name, p.Description)
	}
	if err = iter.Close(); err != nil {
		log.Fatal(err)
	}

	ran := template.Must(template.ParseFiles("reading.html"))
	ran.Execute(w, d)
}

func main() {

	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/reading", reading)
	server.ListenAndServe()

}
