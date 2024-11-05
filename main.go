package main

import (
	"fmt"
	"myChat/repository"
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

	rt := NewRouter(db)

	// サーバ起動設定
	s := http.Server{
		Addr:           conf.Address,
		Handler:        rt,
		ReadTimeout:    time.Duration(conf.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(conf.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
