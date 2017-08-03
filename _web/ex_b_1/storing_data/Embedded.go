package main

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
	"fmt"
)

type Task struct {
	Description string
	Due time.Time
}

type Category struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Description string
	Tasks []Task
}

func main() {
	session, err :=mgo.Dial("localhost")
	if err != nil{
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	C :=session.DB("test").C("mytest")

	doc := Category{
		bson.NewObjectId(),
		"Open-Source",
		"Tasks for open-source projects",
		[]Task{
			Task{"Create project in mgo", time.Date(2015, time.August, 10, 0, 0, 0, 0,
				time.UTC)},
			Task{"Create REST API", time.Date(2015, time.August, 20, 0, 0, 0, 0, time.UTC)},
		},
	}

	err = C.Insert(&doc)
	if err != nil {
		log.Fatal(err)
	}



}