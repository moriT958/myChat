package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Cannot open config file", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal("Cannot get configuration from file", err)
	}
}

func connectDb() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to get db", err)
	}
}
