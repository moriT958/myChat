package controller

import (
	"log"
	"myChat/pkg/utils"
	"net/http"
	"strings"
)

// GET /
// Home page
func (ctlr *Controller) Index(w http.ResponseWriter, req *http.Request) {
	data, err := ctlr.Service.Forum.ReadThreadList()
	if err != nil {
		errMsg := "cannot get threads"
		url := []string{"/err?msg=", errMsg}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
		return
	}

	cookie, err := req.Cookie("_cookie")
	if err != nil {
		log.Println("auth failed, cannot get cookie", err)
		utils.RenderHTML(w, data, "layout", "public.navbar", "index")
		return
	}

	if cookie != nil {
		_, err = ctlr.Service.Auth.CheckSession(cookie.Value)
		if err != nil {
			log.Println(err)
			utils.RenderHTML(w, data, "layout", "public.navbar", "index")
		} else {
			utils.RenderHTML(w, data, "layout", "private.navbar", "index")
		}
	}
}

// GET /err
// error page
func (ctlr *Controller) ErrHandler(w http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()

	cookie, err := req.Cookie("_cookie")
	if err != nil {
		log.Println("auth failed, cannot get cookie", err)
		utils.RenderHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
		return
	}

	_, err = ctlr.Service.Auth.CheckSession(cookie.Value)
	if err != nil {
		utils.RenderHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utils.RenderHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
