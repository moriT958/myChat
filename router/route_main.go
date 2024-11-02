package routers

import (
	"errors"
	"myChat/models"
	"net/http"
	"strings"
)

type Router struct {
	models models.Models
}

func NewRouter(m models.Models) *Router {
	return &Router{models: m}
}

func (r *Router) Index(w http.ResponseWriter, req *http.Request) {
	threads, err := r.models.GetAllThreads()
	if err != nil {
		errMsg := "cannot get threads"
		url := []string{"/err?msg=", errMsg}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
	} else {
		_, err := r.session(w, req)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}

func (r *Router) ErrHandler(w http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	_, err := r.session(w, req)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

// Checks if the user is logged in and has a session, if not err is not nil
func (r *Router) session(_ http.ResponseWriter, req *http.Request) (sess models.Session, err error) {
	cookie, err := req.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(r.models); !ok {
			err = errors.New("invalid session")
		}
	}
	return
}
