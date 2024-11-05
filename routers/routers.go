package routers

import (
	"errors"
	"myChat/repository"
	"net/http"
)

type Router struct {
	repo repository.Repository
}

func NewRouter(repo repository.Repository) *Router {
	return &Router{repo: repo}
}

// Checks if the user is logged in and has a session, if not err is not nil
func (rt *Router) session(req *http.Request) (sess repository.Session, err error) {
	cookie, err := req.Cookie("_cookie")
	if err == nil {
		sess = repository.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("invalid session")
		}
	}
	return
}
