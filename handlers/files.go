package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"go.formulabun.club/site/data"
)

type ServerHandler struct {
	Template *template.Template
	Data     *data.SiteData
}

func (h ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Template.Execute(w, h.Data)
	if err != nil {
		fmt.Println(err)
	}
}
