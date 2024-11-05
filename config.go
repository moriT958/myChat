package main

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var conf *Config

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Cannot open config file", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		log.Fatal("Cannot get configuration from file", err)
	}
}
