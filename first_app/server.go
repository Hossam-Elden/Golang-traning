package main

import (
	"net/http"
	"fmt"
)

func hundler(write http.ResponseWriter , request *http.Request){
	fmt.Fprintf(write ,"Hello world , %s! " , request.URL.Path[1:])
}


func main() {
	http.HandleFunc("/" , hundler)
	http.ListenAndServe(":8080" ,nil)
}
