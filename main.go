package main

import (
	"database/sql"
	"fmt"
	"myChat/models"
	routers "myChat/router"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

const (
	VERSION = "0.1"
)

var (
	config Config
	db     *sql.DB
)

func main() {
	// 設定のロード
	loadConfig()
	fmt.Println("myChat", VERSION, "started at", config.Address)

	// データベースの取得
	connectDb()
	defer db.Close()

	m := models.NewModels(db)  // models層
	r := routers.NewRouter(*m) // routers層

	mux := http.NewServeMux()
	mux.HandleFunc("/", r.Index)         // index
	mux.HandleFunc("/err", r.ErrHandler) // err

	// 静的ファイルの配信
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// サーバ起動設定
	s := http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
