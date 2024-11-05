package controller

import (
	"errors"
	"fmt"
	"html/template"
	"myChat/repository"
	"net/http"
)

// render HTML file responce
func renderHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// Checks if the user is logged in and has a session, if not err is not nil
func session(req *http.Request) (repository.Session, error) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		return repository.Session{}, err
	}

	sess := repository.Session{Uuid: cookie.Value}
	if ok, err := sess.Check(); err != nil {
		return repository.Session{}, err
	} else if !ok {
		return repository.Session{}, errors.New("invalid session")
	}

	return sess, nil
}
