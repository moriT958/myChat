package controller

import (
	"net/http"
	"strings"
)

// GET /
// Home page
func (ctlr *Controller) Index(w http.ResponseWriter, req *http.Request) {
	threads, err := ctlr.repo.GetAllThreads()
	if err != nil {
		errMsg := "cannot get threads"
		url := []string{"/err?msg=", errMsg}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
	} else {
		_, err := session(req)
		if err != nil {
			renderHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			renderHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}

// GET /err
// error page
func (ctlr *Controller) ErrHandler(w http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	_, err := session(req)
	if err != nil {
		renderHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		renderHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
