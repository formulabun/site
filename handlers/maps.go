package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type MapsHandler struct {
	Template *template.Template
}

func (h MapsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  err := h.Template.Execute(w, struct{Tmp string}{"tmp"})
  if err != nil {
    fmt.Println(err)
  }
}
