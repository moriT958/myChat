package routers

import (
	"net/http"
	"strings"
)

func (rt *Router) Index(w http.ResponseWriter, req *http.Request) {
	threads, err := rt.repo.GetAllThreads()
	if err != nil {
		errMsg := "cannot get threads"
		url := []string{"/err?msg=", errMsg}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
	} else {
		_, err := rt.session(req)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}

func (rt *Router) ErrHandler(w http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	_, err := rt.session(req)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
