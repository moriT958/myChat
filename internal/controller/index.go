package controller

import (
	"myChat/pkg/utils"
	"net/http"
	"strings"
)

// GET /
// Home page
func (ctlr *Controller) Index(w http.ResponseWriter, req *http.Request) {
	threads, err := ctlr.tRepo.FindAll()
	if err != nil {
		errMsg := "cannot get threads"
		url := []string{"/err?msg=", errMsg}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
	} else {
		_, err := ctlr.CheckSession(req)
		if err != nil {
			utils.RenderHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			utils.RenderHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}

// GET /err
// error page
func (ctlr *Controller) ErrHandler(w http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	_, err := ctlr.CheckSession(req)
	if err != nil {
		utils.RenderHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utils.RenderHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
