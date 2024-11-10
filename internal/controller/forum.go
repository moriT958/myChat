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
func (ctlr *Controller) ThreadFormHandler(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		log.Println("failed to get cookie: ", err)
		if err == http.ErrNoCookie {
			log.Println(err)
		}
	}
	_, err = ctlr.Service.Auth.CheckSession(cookie.Value)
	if err != nil {
		log.Println("need login to create thread", err)
		http.Redirect(w, req, "/login", http.StatusFound)
	} else {
		utils.RenderHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /thread/create
// Create the user account
func (ctlr *Controller) CreateThreadHandler(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		log.Println("failed to get cookie: ", err)
		if err == http.ErrNoCookie {
			log.Println(err)
		}
	}

	sess, err := ctlr.Service.Auth.CheckSession(cookie.Value)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
	}

	if err := req.ParseForm(); err != nil {
		log.Println(err, "Cannot parse form")
	}

	if err := ctlr.Service.Forum.CreateThread(sess.UserId, req.PostFormValue("topic")); err != nil {
		log.Println("failed to create thread: ", err)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func (ctlr *Controller) ReadThreadHandler(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		log.Println("failed to get cookie: ", err)
		if err == http.ErrNoCookie {
			log.Println(err)
		}
	}

	vals := req.URL.Query()
	uuid := vals.Get("id")

	data, err := ctlr.Service.Forum.ReadThreadDetail(uuid)
	if err != nil {
		http.Redirect(w, req, "/err?msg=Cannot read thread", http.StatusFound)
	}

	// Check session, and Navigate page.
	if _, err := ctlr.Service.Auth.CheckSession(cookie.Value); err != nil {
		utils.RenderHTML(w, data, "layout", "public.navbar", "public.thread")
	} else {
		utils.RenderHTML(w, data, "layout", "private.navbar", "private.thread")
	}
}

// POST /thread/post
// Create the post
func (ctlr *Controller) PostThreadHandler(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		log.Println("failed to get cookie: ", err)
		if err == http.ErrNoCookie {
			log.Println(err)
		}
	}

	sess, err := ctlr.Service.Auth.CheckSession(cookie.Value)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
	}
	if err := req.ParseForm(); err != nil {
		log.Println(err, "Cannot parse form")
	}

	body := req.PostFormValue("body")
	threadUuid := req.PostFormValue("uuid")
	err = ctlr.Service.Forum.CreatePost(sess.UserId, body, threadUuid)
	if err != nil {
		url := []string{"/err?msg=", "Cannot read thread"}
		http.Redirect(w, req, strings.Join(url, ""), http.StatusFound)
	}

	url := fmt.Sprint("/thread/read?id=", threadUuid)
	http.Redirect(w, req, url, http.StatusFound)
}
