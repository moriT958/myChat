package postgres

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once sync.Once
	db   *sql.DB
)

func Connect() *sql.DB {
	once.Do(func() {
		var err error
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("failed load dotenv file: ", err)
		}
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
