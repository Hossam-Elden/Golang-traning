package main

import (
	"net/http"
	"encoding/base64"
	"fmt"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request){
	msg:=[]byte("hello motherfathersssss")
	f:=http.Cookie{
		Name:"flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &f)
}


func showMessage(w http.ResponseWriter, r *http.Request) {
	f, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name: "flash",
			MaxAge: -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(f.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)

	server.ListenAndServe()
}