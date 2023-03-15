package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandrel())
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("Can not create server [stop]")
	}
}

func helloHandrel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	}
}
