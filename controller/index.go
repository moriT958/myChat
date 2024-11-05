package controller

import (
	"net/http"
	"strings"
)

func (ctlr *Controller) Index(w http.ResponseWriter, req *http.Request) {
	threads, err := ctlr.repo.GetAllThreads()
	if err != nil {
		errMsg := "cannot get threads"
		url := []string{"/err?msg=", errMsg}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
	} else {
		_, err := session(req)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}

func (ctlr *Controller) ErrHandler(w http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	_, err := session(req)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
