package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"go.formulabun.club/site/handlers"
)

func main() {

	http.Handle("/", http.RedirectHandler("/about", http.StatusFound))
	http.Handle("/public/", http.FileServer(http.Dir(".")))

	fs := os.DirFS("templates/")
  baseTemplate := template.Must(template.ParseFS(fs, "base.tmpl"))

	about := handlers.AboutHandler{
    template.Must(template.Must(baseTemplate.Clone()).ParseFS(fs, "about.tmpl")),
	}
	http.Handle("/about", about)

	maps := handlers.MapsHandler{
    template.Must(template.Must(baseTemplate.Clone()).ParseFS(fs, "maps.tmpl")),
	}
	http.Handle("/maps", maps)

	log.Println("serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
