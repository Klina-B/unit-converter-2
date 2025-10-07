package handlers

import (
	"html/template"
	"net/http"
)

func WeightHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		tmpl := template.Must(template.ParseFiles("templates/weight.html"))
		tmpl.Execute(w, nil)
	}
}