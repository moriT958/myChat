package controller

import (
	"log"
	"myChat/pkg/apperrors"
	"myChat/pkg/utils"
	"net/http"
)

// GET /signup
// Show the signup page
func (ctlr *Controller) SignupFormHandler(w http.ResponseWriter, r *http.Request) {
	err := utils.RenderHTML(w, nil, "login.layout", "public.navbar", "signup")
	if err != nil {
		err = apperrors.RenderHTMLFailed.Wrap(err, "failed render html")
		apperrors.ErrorHandler(w, r, err)
	}
}

// POST /signup_account
// Create the user account
func (ctlr *Controller) SignupPostHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err, "Cannot parse form")
	}
	err = ctlr.Service.Auth.CreateUser(
		req.PostFormValue("name"),
		req.PostFormValue("email"),
		req.PostFormValue("password"),
	)
	if err != nil {
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
	uuid, err := ctlr.Service.Auth.Login(
		req.PostFormValue("email"),
		req.PostFormValue("password"),
	)
	if err != nil {
		log.Println("failed to login: ", err)
		http.Redirect(w, req, "/login", http.StatusFound)
	}

	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, req, "/", http.StatusFound)
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
	if err := ctlr.Service.Auth.Logout(cookie.Value); err != nil {
		log.Println("failed to logout: ", err)
	}

	http.Redirect(w, req, "/", http.StatusFound)
}
