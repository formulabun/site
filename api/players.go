package api

import (
	"net/http"
	"text/template"

	"go.formulabun.club/site/data"
)

type playerHandler struct {
	template *template.Template
	data     *[]data.Player
}

func newPlayerHandler(tmpl *template.Template, data *[]data.Player) *playerHandler {
	return &playerHandler{
		tmpl,
		data,
	}
}

func (a *playerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if a.data != nil {
		a.template.Execute(w, a.data)
	}
}
