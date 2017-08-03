package main

import (
	"net/http"
	"fmt"
)

func server(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World omg it's build, %s!", r.URL.Path[1:])
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", server)
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}