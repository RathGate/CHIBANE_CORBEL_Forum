package main

import (
	"html/template"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	// *Generates and executes templates:
	tmpl := template.Must(template.ParseFiles("templates/views/index.html"))
	tmpl.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// *Generates and executes templates:
	tmpl := template.Must(template.ParseFiles("templates/views/index.html"))
	tmpl.Execute(w, nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/views/register_form.html"))
	tmpl.Execute(w, nil)
}
