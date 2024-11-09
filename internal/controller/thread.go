package controller

import (
	"fmt"
	"log"
	"myChat/internal/model"
	"myChat/pkg/utils"
	"net/http"
	"strings"
	"time"
)

// GET /threads/new
// Show the new thread form page
func (ctlr Controller) ThreadFormHandler(w http.ResponseWriter, req *http.Request) {
	_, err := ctlr.CheckSession(req)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
	} else {
		utils.RenderHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /thread/create
// Create the user account
func (ctlr Controller) CreateThreadHandler(w http.ResponseWriter, req *http.Request) {
	sess, err := ctlr.CheckSession(req)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
	}

	if err := req.ParseForm(); err != nil {
		log.Println(err, "Cannot parse form")
	}

	user, err := ctlr.uRepo.FindById(sess.UserId)
	if err != nil {
		log.Println(err, "Cannot get user from utils.Session")
	}

	thread := model.Thread{
		Uuid:      utils.CreateUUID(),
		Topic:     req.PostFormValue("topic"),
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	if err := ctlr.tRepo.Save(thread); err != nil {
		log.Println(err, "Cannot create thread")
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func (ctlr Controller) ReadThreadHandler(w http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	uuid := vals.Get("id")
	thread, err := ctlr.tRepo.FindByUuid(uuid)
	if err != nil {
		url := []string{"/err?msg=", "Cannot read thread"}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
	}

	if _, err := ctlr.CheckSession(req); err != nil {
		utils.RenderHTML(w, &thread, "layout", "public.navbar", "public.thread")
	} else {
		utils.RenderHTML(w, &thread, "layout", "private.navbar", "private.thread")
	}
}

// POST /thread/post
// Create the post
func (ctlr Controller) PostThreadHandler(w http.ResponseWriter, req *http.Request) {
	sess, err := ctlr.CheckSession(req)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
	}
	if err := req.ParseForm(); err != nil {
		log.Println(err, "Cannot parse form")
	}
	user, err := ctlr.uRepo.FindById(sess.UserId)
	if err != nil {
		log.Println(err, "Cannot get user from session")
	}

	body := req.PostFormValue("body")
	uuid := req.PostFormValue("uuid")
	thread, err := ctlr.tRepo.FindByUuid(uuid)
	if err != nil {
		url := []string{"/err?msg=", "Cannot read thread"}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
	}

	post := model.Post{
		Uuid:      uuid,
		Body:      body,
		UserId:    user.Id,
		ThreadId:  thread.Id,
		CreatedAt: time.Now(),
	}
	if err := ctlr.pRepo.Save(post); err != nil {
		log.Println(err, "Cannot create post")
	}
	url := fmt.Sprint("/thread/read?id=", uuid)
	http.Redirect(w, req, url, http.StatusFound)
}
