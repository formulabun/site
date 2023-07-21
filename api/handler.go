package api

import (
	"context"
	"text/template"

	"github.com/gorilla/mux"
	"go.formulabun.club/site/data"
)

type ApiHandler struct {
	data *data.SiteData
}

func Route(r *mux.Router) {
	d := data.Init(context.Background())

	playersTmpl := template.Must(template.ParseFiles("templates/players.tmpl"))
	r.Handle("/players", newPlayerHandler(playersTmpl, &d.Players))

	serverTmpl := template.Must(template.ParseFiles("templates/server.tmpl"))
	r.Handle("/server", newServerHandler(serverTmpl, &d.ServerInfo))

	tmpl := template.Must(template.ParseFiles("templates/files.tmpl"))
	r.Handle("/files/maps", newFilesHandler(tmpl, &d.Maps))
	r.Handle("/files/chars", newFilesHandler(tmpl, &d.Characters))
	r.Handle("/files/mods", newFilesHandler(tmpl, &d.Mods))
	r.Handle("/files/other", newFilesHandler(tmpl, &d.Other))
}
