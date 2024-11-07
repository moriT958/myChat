package main

import (
	"fmt"
	"myChat/api"
	"myChat/config"
	"myChat/internal/repository"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

const (
	VERSION = "0.1"
)

func main() {
	// 設定のロード
	cfg := config.Load("config.json")
	fmt.Println("myChat", VERSION, "started at", cfg.Address)

	// データベースの取得
	db := repository.GetDB()
	defer db.Close()

	rt := api.NewRouter(db)

	// サーバ起動設定
	s := http.Server{
		Addr:           cfg.Address,
		Handler:        rt,
		ReadTimeout:    time.Duration(cfg.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(cfg.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
