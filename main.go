package main

import (
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	port := os.GetEnv("PORT")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}