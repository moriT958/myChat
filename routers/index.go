package routers

import (
	"myChat/models"
	"net/http"
	"strings"
)

func (r *Router) Index(w http.ResponseWriter, req *http.Request) {
	threads, err := r.models.GetAllThreads()
	data := struct {
		Threads []models.Thread
		Models  models.Models
	}{
		Threads: threads,
		Models:  r.models,
	}
	if err != nil {
		errMsg := "cannot get threads"
		url := []string{"/err?msg=", errMsg}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
	} else {
		_, err := r.session(req)
		if err != nil {
			generateHTML(w, data, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, data, "layout", "private.navbar", "index")
		}
	}
}

func (r *Router) ErrHandler(w http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	_, err := r.session(req)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
