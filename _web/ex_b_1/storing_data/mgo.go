package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	//Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//get collection
	c := session.DB("test").C("test")

	docM := map[string]string{
		"name":        "Open Sourcesdasdasdasdasdasd",
		"description": "Tasks for open-source projects",
	}

	//insert a map object
	err = c.Insert(docM)
	if err != nil {
		log.Fatal(err)
	}
	docD := bson.D{
		{"name", "Project"},
		{"description", "Project Tasks"},
	}





	//insert a document slice
	err = c.Insert(docD)
	if err != nil {
		log.Fatal(err)
	}




}
