package controller

import (
	"errors"
	"log"
	"myChat/internal/model"
	"myChat/internal/repository"
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
		Name:     req.PostFormValue("name"),
		Email:    req.PostFormValue("email"),
		Password: req.PostFormValue("password"),
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
	req.ParseForm()
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
		// TODO:
		// 現状user情報を取得する方法がない。
		// そして、sessionがuserエンティティからしか取得できないため、sessionを削除できない。
		// sessionとuserを分けて別々のリポジトリを作るべきか？
		session := repository.Session{Uuid: cookie.Value}
		session.Delete()
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

// Checks if the user is logged in and has a session, if not err is not nil
func (ctlr Controller) CheckSession(req *http.Request) (repository.Session, error) {
	cookie, err := req.Cookie("_cookie")
	if err != nil {
		return repository.Session{}, err
	}

	sess := repository.Session{Uuid: cookie.Value}
	if ok, err := sess.Check(); err != nil {
		return repository.Session{}, err
	} else if !ok {
		return repository.Session{}, errors.New("invalid session")
	}

	return sess, nil
}
