package postgres

import (
	"database/sql"
	"log"
	"os"
	"sync"
)

var (
	once sync.Once
	db   *sql.DB
)

func Connect() *sql.DB {
	once.Do(func() {
		var err error
		dsn := os.Getenv("DATABASE_URL")
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("failed to get db conn: %v", err)
		}

		// Check db responce
		if err = db.Ping(); err != nil {
			log.Fatalf("db response failed: %v", err)
		}
	})
	return db
}
