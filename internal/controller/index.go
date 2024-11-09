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
	}

	data := make([]struct {
		Uuid       string
		Topic      string
		UserName   string
		CreatedAt  string
		NumReplies int
	}, len(threads))

	for i, thread := range threads {
		data[i].Uuid = thread.Uuid
		data[i].Topic = thread.Topic
		usr, _ := ctlr.uRepo.FindById(thread.Id)
		data[i].UserName = usr.Name
		data[i].CreatedAt = thread.CreatedAtStr()
		repNum, _ := ctlr.tRepo.CountPostNum(thread.Id)
		data[i].NumReplies = repNum
	}

	_, err = ctlr.CheckSession(req)
	if err != nil {
		utils.RenderHTML(w, data, "layout", "public.navbar", "index")
	} else {
		utils.RenderHTML(w, data, "layout", "private.navbar", "index")
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
