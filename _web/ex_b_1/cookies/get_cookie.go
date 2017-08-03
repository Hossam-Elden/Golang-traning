package main

import (
	"net/http"
	"fmt"
)

func setCookie(w http.ResponseWriter, r *http.Request){
	k1 := http.Cookie{
		Name: "one_cookie",
		Value: "Go Web Programming",
		HttpOnly: true,
	}
	k2 :=http.Cookie{
		Name: "second_cookie",
		Value: "Go Web Programming 2 ",
		HttpOnly: true,
	}
	http.SetCookie(w, &k1)
	http.SetCookie(w, &k2)
}

func getCookie(w http.ResponseWriter, r *http.Request){
	k1, err := r.Cookie("one_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, k1)
	fmt.Fprintln(w, cs)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)

	server.ListenAndServe()
}