package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		tmpl := template.Must(template.ParseFiles("templates/length.html"))
		tmpl.Execute(w, nil)
	} else{
		value, _ := strconv.ParseFloat(r.FormValue("lengthValue"), 64)
		fromUnit := r.FormValue("fromUnit")
		toUnit := r.FormValue("toUnit")
		result := UnitLength{Value: value, Unit: fromUnit}.Convert(toUnit)

		tmpl := template.Must(template.ParseFiles("templates/length.html"))
		tmpl.Execute(w, map[string]any{
			"Result": result,
			"ToUnit": toUnit,
		})
	}
}