package controller

import (
	"log"
	"myChat/internal/domain/model"
	"net/http"
)

// find session in database
// Checks if the user is logged in and has a session, if not err is not nil
func (ctlr *Controller) CheckSession(req *http.Request) (model.Session, error) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		log.Println("auth failed, cannot get cookie", err)
		return model.Session{}, err
	}

	// return empty session and error,
	// if session doesnt exits in database.
	session, err := ctlr.sRepo.FindByUuid(cookie.Value)
	if (err != nil || session == model.Session{}) {
		log.Println("auth failed, cannot get user session", err)
		return model.Session{}, err
	}

	return session, nil
}
