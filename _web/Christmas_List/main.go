package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"github.com/martini-contrib/binding"
)


type Wish struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}


// DB Returns a martini.Handler
func DB() martini.Handler {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB("advent"))
		defer s.Close()
		c.Next()
	}
}

// GetAll returns all Wishes in the database
func GetAll(db *mgo.Database) []Wish {
	var wishlist []Wish
	db.C("wishes").Find(nil).All(&wishlist)
	return wishlist
}


func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	//START3 OMIT
	m.Use(DB())

	m.Get("/wishes", func(r render.Render, db *mgo.Database) {
		r.HTML(200, "list", GetAll(db))
	})

	m.Post("/wishes", binding.Form(Wish{}), func(wish Wish, r render.Render, db *mgo.Database) {
		db.C("wishes").Insert(wish)
		r.HTML(200, "list", GetAll(db))
	})

	m.Run()
}