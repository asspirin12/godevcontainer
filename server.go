package main

import (
	"io"
	"net/http"

	"hello/user"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, user.Hello())
}

func main() {
	http.HandleFunc("/hello", handle)
	http.ListenAndServe(":8080", nil)
}
