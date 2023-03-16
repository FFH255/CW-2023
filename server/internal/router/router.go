package router

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Configure() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/hello", handleHello())

	return router
}

func handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world")
	}
}
