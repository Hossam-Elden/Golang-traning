package main

import (
"text/template"
"os"
"log"
)

//pointer for the template file from the type template
var tpl *template.Template

//initialise function before the main function
func init() {

	//must will exit if there is err if not will create the template from the index file
	tpl = template.Must(template.ParseFiles("texts/list.hos"))
}

func main() {
	//if there this no err will execute os func and send 20 as data for the index file
	err := tpl.Execute(os.Stdout, []string{"one","tow","there"})
	if err != nil {
		log.Fatalln(err)
	}
}

