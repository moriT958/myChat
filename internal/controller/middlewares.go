package controller

import (
	"myChat/internal/model"
	"net/http"
)

// find session in database
// Checks if the user is logged in and has a session, if not err is not nil
func (ctlr *Controller) CheckSession(req *http.Request) (model.Session, error) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		return model.Session{}, err
	}

	// return empty session and error,
	// if session doesnt exits in database.
	session, err := ctlr.sRepo.FindByUuid(cookie.Value)
	if (err != nil || session == model.Session{}) {
		return model.Session{}, err
	}

	return session, nil
}
