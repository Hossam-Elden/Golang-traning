package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"os"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	Col := session.DB("test").C("mytest")
	// Change the winner for game 55da80 to Seth
	gameId := bson.ObjectIdHex("597cf7a76f71d7136422a190")
	newWinner := "Seth"
	update := bson.M{"$set": bson.M{"name": newWinner}}
	if err := Col.UpdateId(gameId, update); err != nil {
		panic(err)
		os.Exit(1)

	}

	fmt.Printf("name of data %s updated to %s.\n", gameId, newWinner)
}
