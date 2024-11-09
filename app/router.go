package app

import (
	"database/sql"
	"myChat/internal/controller"
	"myChat/internal/domain/repository"
	"net/http"
)

func NewRouter(db *sql.DB) (mux *http.ServeMux) {
	// object dependencies
	// repositories depends on db: *sql.DB
	// service depends on repository
	// controller depends on service
	// mux depends on controller
	uRepo := repository.NewUserRepository(db)
	sRepo := repository.NewSessionRepository(db)
	tRepo := repository.NewThreadRepository(db)
	pRepo := repository.NewPostRepository(db)
	// ser := service.NewAppService(*uRepo, *tRepo)
	ctlr := controller.NewController(*uRepo, *sRepo, *tRepo, *pRepo)

	mux = http.NewServeMux()
	mux.HandleFunc("GET /", ctlr.Index)         // index
	mux.HandleFunc("GET /err", ctlr.ErrHandler) // err

	// Defined in controller directory
	// authentication handlers defined in auth.go
	mux.HandleFunc("GET /login", ctlr.LoginFormHandler)
	mux.HandleFunc("GET /logout", ctlr.LogoutHandler)
	mux.HandleFunc("GET /signup", ctlr.SignupFormHandler)
	mux.HandleFunc("POST /signup_account", ctlr.SignupPostHandler)
	mux.HandleFunc("POST /authenticate", ctlr.AuthenticateHandler)

	// // thread handlers difined in thread.go
	mux.HandleFunc("GET /threads/new", ctlr.ThreadFormHandler)
	mux.HandleFunc("POST /thread/create", ctlr.CreateThreadHandler)
	mux.HandleFunc("POST /thread/post", ctlr.PostThreadHandler)
	mux.HandleFunc("GET /thread/read", ctlr.ReadThreadHandler)

	// Serves static contents
	files := http.FileServer(http.Dir("web"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", files))

	return
}
