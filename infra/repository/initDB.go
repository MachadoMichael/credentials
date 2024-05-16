package repository

import (
	"database/sql"
	"log"
)

// var DB *sql.DB
var AuthRepo *BasicAuthRepo

func InitDB() {
	var err error
	DB, err := sql.Open("sqlite3", "../db/")

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Database ping failed %v", err)
	}

	AuthRepo = NewBasicAuthRepo(DB)

	log.Println("Successfully connected to SQLite database.")
}
