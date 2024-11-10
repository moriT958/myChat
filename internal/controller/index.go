package controller

import (
	"log"
	"myChat/pkg/utils"
	"net/http"
	"strings"
)

type indexPageData struct {
	Threads []struct {
		Topic      string
		Uuid       string
		UserName   string
		CreatedAt  string
		NumReplies int
	}
}

// GET /
// Home page
func (ctlr *Controller) Index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		log.Println("auth failed, cannot get cookie", err)
	}

	data, err := ctlr.Service.Forum.ReadThreadList()
	if err != nil {
		errMsg := "cannot get threads"
		url := []string{"/err?msg=", errMsg}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
	}

	_, err = ctlr.Service.Auth.CheckSession(cookie.Value)
	if err != nil {
		utils.RenderHTML(w, data, "layout", "public.navbar", "index")
	} else {
		utils.RenderHTML(w, data, "layout", "private.navbar", "index")
	}
}

// GET /err
// error page
func (ctlr *Controller) ErrHandler(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		log.Println("auth failed, cannot get cookie", err)
	}

	vals := req.URL.Query()
	_, err = ctlr.Service.Auth.CheckSession(cookie.Value)
	if err != nil {
		utils.RenderHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utils.RenderHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
