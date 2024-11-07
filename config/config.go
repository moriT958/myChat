package config

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Version      string
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

func Load(filename string) (cfg *Config) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Cannot open config file", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		log.Fatal("Cannot get configuration from file", err)
	}
	return
}
