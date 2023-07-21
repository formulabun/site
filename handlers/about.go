package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"go.formulabun.club/site/data"
)

type AboutHandler struct {
	Template *template.Template
	Data     *data.SiteData
}

func (h AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Template.Execute(w, h.Data)
	if err != nil {
		fmt.Println(err)
	}
}
