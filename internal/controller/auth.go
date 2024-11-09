package controller

import (
	"log"
	"myChat/internal/model"
	"myChat/pkg/utils"
	"net/http"
	"time"
)

// GET /signup
// Show the signup page
func (ctlr *Controller) SignupFormHandler(w http.ResponseWriter, _ *http.Request) {
	utils.RenderHTML(w, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup_account
// Create the user account
func (ctlr *Controller) SignupPostHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err, "Cannot parse form")
	}
	user := model.User{
		Uuid:      utils.CreateUUID(),
		Name:      req.PostFormValue("name"),
		Email:     req.PostFormValue("email"),
		Password:  utils.Encrypt(req.PostFormValue("password")),
		CreatedAt: time.Now(),
	}
	if err := ctlr.uRepo.Save(user); err != nil {
		log.Println(err, "Cannot create user")
	}
	http.Redirect(w, req, "/login", http.StatusFound)
}

// GET /login
// login form page
func (ctlr *Controller) LoginFormHandler(w http.ResponseWriter, _ *http.Request) {
	utils.RenderHTML(w, nil, "login.layout", "public.navbar", "login")
}

// POST /authenticate
// Authenticate the user given the email and password
func (ctlr *Controller) AuthenticateHandler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Println("cannot parse form: ", err)
	}
	user, err := ctlr.uRepo.FindByEmail(req.PostFormValue("email"))
	if err != nil {
		log.Println(err, "Cannot find user by email")
	}
	if user.Password == utils.Encrypt(req.PostFormValue("password")) {
		session := model.Session{
			Uuid:      utils.CreateUUID(),
			Email:     user.Email,
			UserId:    user.Id,
			CreatedAt: time.Now(),
		}
		if err := ctlr.sRepo.Save(session); err != nil {
			log.Println(err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, req, "/", http.StatusFound)
	} else {
		http.Redirect(w, req, "/login", http.StatusFound)
	}

}

// GET /logout
// Logs the user out
func (ctlr *Controller) LogoutHandler(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		log.Println("failed to get cookie: ", err)
		if err == http.ErrNoCookie {
			log.Println(err)
		}
	}

	// delete all user's session
	session, err := ctlr.sRepo.FindByUuid(cookie.Value)
	if err != nil {
		log.Println(err)
	}
	if err := ctlr.sRepo.DeleteByUserId(session.UserId); err != nil {
		log.Println(err)
	}

	http.Redirect(w, req, "/", http.StatusFound)
}
