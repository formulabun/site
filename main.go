package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.formulabun.club/site/api"
)

func main() {
	r := mux.NewRouter()

	// Order of matchers is important!
	assetHandler := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/public/").Handler(assetHandler)

	apiR := r.PathPrefix("/api").Subrouter()
	api.Route(apiR)

	htmlHandler := http.FileServer(http.Dir("static/html/"))
	r.PathPrefix("/").Handler(htmlHandler)

	http.ListenAndServe(":8080", r)
}
