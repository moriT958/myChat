package main

import (
	"database/sql"
	"myChat/controller"
	"myChat/repository"
	"net/http"
)

func NewRouter(db *sql.DB) (mux *http.ServeMux) {
	repo := repository.NewRepository(db)
	ctlr := controller.NewController(*repo)

	mux = http.NewServeMux()
	mux.HandleFunc("/", ctlr.Index)         // index
	mux.HandleFunc("/err", ctlr.ErrHandler) // err
	mux.HandleFunc("/login", ctlr.LoginHandler)

	// 静的ファイルの配信
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	return
}
