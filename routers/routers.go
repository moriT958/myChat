package routers

import (
	"errors"
	"myChat/models"
	"net/http"
)

type Router struct {
	models models.Models
}

func NewRouter(m models.Models) *Router {
	return &Router{models: m}
}

// Checks if the user is logged in and has a session, if not err is not nil
func (r *Router) session(req *http.Request) (sess models.Session, err error) {
	cookie, err := req.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(r.models); !ok {
			err = errors.New("invalid session")
		}
	}
	return
}
