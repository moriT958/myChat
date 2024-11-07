package controller

import (
	"fmt"
	"log"
	"myChat/pkg/utils"
	"net/http"
	"strings"
)

// GET /threads/new
// Show the new thread form page
func (ctlr Controller) ThreadFormHandler(w http.ResponseWriter, req *http.Request) {
	_, err := utils.Session(req)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
	} else {
		utils.RenderHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /thread/create
// Create the user account
func (ctlr Controller) CreateThreadHandler(w http.ResponseWriter, req *http.Request) {
	sess, err := utils.Session(req)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
	} else {
		err = req.ParseForm()
		if err != nil {
			log.Println(err, "Cannot parse form")
		}
		user, err := sess.GetUser()
		if err != nil {
			log.Println(err, "Cannot get user from utils.Session")
		}
		topic := req.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			log.Println(err, "Cannot create thread")
		}
		http.Redirect(w, req, "/", http.StatusFound)
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func (ctlr Controller) ReadThreadHandler(w http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	uuid := vals.Get("id")
	thread, err := ctlr.repo.GetThreadByUUID(uuid)
	if err != nil {
		url := []string{"/err?msg=", "Cannot read thread"}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
	} else {
		_, err := utils.Session(req)
		if err != nil {
			utils.RenderHTML(w, &thread, "layout", "public.navbar", "public.thread")
		} else {
			utils.RenderHTML(w, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
// Create the post
func (ctlr Controller) PostThreadHandler(w http.ResponseWriter, req *http.Request) {
	sess, err := utils.Session(req)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
	} else {
		err = req.ParseForm()
		if err != nil {
			log.Println(err, "Cannot parse form")
		}
		user, err := sess.GetUser()
		if err != nil {
			log.Println(err, "Cannot get user from session")
		}
		body := req.PostFormValue("body")
		uuid := req.PostFormValue("uuid")
		thread, err := ctlr.repo.GetThreadByUUID(uuid)
		if err != nil {
			url := []string{"/err?msg=", "Cannot read thread"}
			http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			log.Println(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(w, req, url, http.StatusFound)
	}
}
