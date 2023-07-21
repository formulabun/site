package api

import (
	"net/http"
	"text/template"
)

type filesHandler struct {
	template *template.Template
	files    *[]string
}

func newFilesHandler(tmpl *template.Template, data *[]string) *filesHandler {
	return &filesHandler{
		tmpl,
		data,
	}
}

func (a *filesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if a.files != nil {
		a.template.Execute(w, a.files)
	}
}
