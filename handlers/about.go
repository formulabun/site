package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type AboutHandler struct {
	Template *template.Template
}

func (h AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  err := h.Template.Execute(w, struct{}{})
  if err != nil {
    fmt.Println(err)
  }
}
