package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	connStr := os.Getenv("POSTGRES_CONN_STRING")
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	log.Println("Connected to PostgreSQL!")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
