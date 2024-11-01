package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	version      = "0.1"
	adress       = "0.0.0.0:8080"
	readTimeout  = int64(10)
	writeTimeout = int64(600)
)

func main() {
	fmt.Println("myChat", version, "started at", adress)

	m := http.NewServeMux()
	// 静的ファイルの配信
	files := http.FileServer(http.Dir("public"))
	m.Handle("/static/", http.StripPrefix("/static/", files))

	m.HandleFunc("/", index)  // index
	m.HandleFunc("/err", err) // err

	// サーバ起動設定
	s := http.Server{
		Addr:           adress,
		Handler:        m,
		ReadTimeout:    time.Duration(readTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(writeTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
