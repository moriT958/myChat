package main

import (
	"fmt"
	"myChat/app"
	"myChat/config"
	"myChat/pkg/postgres"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Load configurations
	cfg := config.Load("config.json")
	fmt.Printf("mychat App(version%s) Started on %s\n", cfg.Version, cfg.Address)

	// Connect DB
	db := postgres.Connect()
	defer db.Close()

	hdlr := app.NewAppHandler(db)

	// server settings
	s := http.Server{
		Addr:           cfg.Address,
		Handler:        hdlr,
		ReadTimeout:    time.Duration(cfg.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(cfg.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
