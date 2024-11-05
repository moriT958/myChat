package controller

import (
	"html/template"
	"net/http"
)

// GET /login
// Show the login page
func (ctlr *Controller) LoginHandler(w http.ResponseWriter, _ *http.Request) {
	files := []string{"login.layout", "public.navbar", "login"}
	for i := range files {
		files[i] = "templates/" + files[i] + ".html"
	}

	t := template.New("layout")
	t = template.Must(t.ParseFiles(files...))
	t.Execute(w, nil)
}
