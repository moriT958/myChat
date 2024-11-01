package main

import (
	"database/sql"
	"fmt"
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

	m := http.NewServeMux()
	// 静的ファイルの配信
	files := http.FileServer(http.Dir("public"))
	m.Handle("/static/", http.StripPrefix("/static/", files))

	m.HandleFunc("/", index)         // index
	m.HandleFunc("/err", errHandler) // err

	// サーバ起動設定
	s := http.Server{
		Addr:           config.Address,
		Handler:        m,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
