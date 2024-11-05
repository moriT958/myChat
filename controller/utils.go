package controller

import (
	"errors"
	"fmt"
	"html/template"
	"myChat/repository"
	"net/http"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// Checks if the user is logged in and has a session, if not err is not nil
func session(req *http.Request) (sess repository.Session, err error) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		return
	}

	sess = repository.Session{Uuid: cookie.Value}
	if ok, _ := sess.Check(); !ok {
		err = errors.New("invalid session")
	}

	return
}
