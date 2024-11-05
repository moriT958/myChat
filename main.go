package main

import (
	"fmt"
	"myChat/repository"
	"myChat/routers"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

const (
	VERSION = "0.1"
)

func main() {
	// 設定のロード
	loadConfig()
	fmt.Println("myChat", VERSION, "started at", conf.Address)

	// データベースの取得
	db := repository.GetDB()
	defer db.Close()

	repo := repository.NewRepository(db)
	r := routers.NewRouter(*repo) // routers層

	mux := http.NewServeMux()
	mux.HandleFunc("/", r.Index)         // index
	mux.HandleFunc("/err", r.ErrHandler) // err

	// 静的ファイルの配信
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// サーバ起動設定
	s := http.Server{
		Addr:           conf.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(conf.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(conf.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
