package main

import (
	"html/template"
	"log"
	"myChat/models"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/index.html",
	}

	m := models.NewModels(db)

	users, _ := m.GetAllUsers()
	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w, "layout", users)
	if err != nil {
		log.Println((err))
	}
}

func errHandler(w http.ResponseWriter, r *http.Request) {

}
