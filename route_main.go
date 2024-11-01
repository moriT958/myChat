package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Println((err))
	}
}

func errHandler(w http.ResponseWriter, r *http.Request) {

}
