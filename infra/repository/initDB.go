package repository

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
)

// var db *sql.DB
var AuthRepo *BasicAuthRepo

func InitDB() {
	var err error
	dbPath := "../db/"
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database ping failed %v", err)
	}

	AuthRepo = NewBasicAuthRepo(db)

	log.Println("Successfully connected to SQLite database.")

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChan
		closeDB()
		os.Exit(0)
	}()
}

func closeDB() {
	if AuthRepo != nil && AuthRepo.db != nil {
		_ = AuthRepo.db.Close()
		log.Println("Database connection closed.")
	}
}
