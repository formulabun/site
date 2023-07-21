package api

import (
	"net/http"
	"text/template"

	"go.formulabun.club/site/data"
)

type serverHandler struct {
	template *template.Template
	data     **data.ServerInfo
}

func newServerHandler(tmpl *template.Template, data **data.ServerInfo) *serverHandler {
	return &serverHandler{
		tmpl,
		data,
	}
}

func (a *serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if *a.data != nil {
		a.template.Execute(w, a.data)
	}
}
