package main

import (
	"text/template"
	"log"
	"os"
)

func main()  {
	tpl, err := template.ParseFiles("index.hos")
	if err != nil{
		log.Fatalln(err)
	}


	err=tpl.Execute(os.Stdout, nil)
	if err != nil{
		log.Fatalln(err)
	}

}
