package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"os"
)

func main() {
	session, err:=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	defer  session.Close()
	session.SetMode(mgo.Monotonic,true)

	Col:=session.DB("test").C("mytest")

	// Delete record
	err = Col.Remove(bson.M{"name": "Open-Source"})

	if err != nil {
		fmt.Printf("remove fail %v\n", err)
		os.Exit(1)
	}
}
