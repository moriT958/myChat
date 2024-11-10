package app

import (
	"database/sql"
	"myChat/internal/controller"
	"myChat/internal/domain/repository"
	"myChat/internal/service"
	"net/http"
)

func NewAppHandler(db *sql.DB) (mux *http.ServeMux) {

	// di
	ctlr := resolveDependencyToController(db)

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

// object dependencies
// repositories depends on db: *sql.DB
// service depends on repository
// controller depends on service

func resolveDependencyToController(db *sql.DB) *controller.Controller {
	uRepo := repository.NewUserRepository(db)
	sRepo := repository.NewSessionRepository(db)
	tRepo := repository.NewThreadRepository(db)
	pRepo := repository.NewPostRepository(db)

	aSer := service.NewAuthService(*uRepo, *sRepo)
	fSer := service.NewForumService(*tRepo, *pRepo, *uRepo)
	ser := service.NewAppService(aSer, fSer)

	ctlr := controller.NewController(ser)
	return ctlr
}
