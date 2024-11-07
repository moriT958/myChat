package controller

import (
	"log"
	"myChat/internal/repository"
	"myChat/pkg/utils"
	"net/http"
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
	user := repository.User{
		Name:     req.PostFormValue("name"),
		Email:    req.PostFormValue("email"),
		Password: req.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
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
	req.ParseForm()
	user, err := ctlr.repo.GetUserByEmail(req.PostFormValue("email"))
	if err != nil {
		log.Println(err, "Cannot find user")
	}
	if user.Password == repository.Encrypt(req.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err, "Cannot create session")
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
	if err != http.ErrNoCookie {
		log.Println(err, "Failed to get cookie")
		session := repository.Session{Uuid: cookie.Value}
		session.Delete()
	}
	http.Redirect(w, req, "/", http.StatusFound)
}
