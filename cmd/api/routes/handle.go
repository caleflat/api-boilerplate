package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handle(router *mux.Router) {
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World"))
	})
}
