package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"

	"go.formulabun.club/site/data"
	"go.formulabun.club/site/handlers"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	d := data.Init(ctx)

	http.Handle("/", http.RedirectHandler("/about", http.StatusFound))
	http.Handle("/public/", http.FileServer(http.Dir(".")))

	fs := os.DirFS("templates/")
	baseTemplate := template.Must(template.ParseFS(fs, "base.tmpl"))

	about := handlers.AboutHandler{
		template.Must(template.Must(baseTemplate.Clone()).ParseFS(fs, "about.tmpl")),
		d,
	}
	http.Handle("/about", about)

	maps := handlers.ServerHandler{
		template.Must(template.Must(baseTemplate.Clone()).ParseFS(fs, "files.tmpl")),
		d,
	}
	http.Handle("/files", maps)

	log.Println("serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	cancel()
}
