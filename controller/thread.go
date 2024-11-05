package controller

import (
	"net/http"
)

// GET /threads/new
// Show the new thread form page
func ThreadFormHandler(w http.ResponseWriter, req *http.Request) {
	_, err := session(req)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}
