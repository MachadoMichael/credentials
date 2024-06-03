package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var CredentialRepo *Repo
var client *redis.Client

func Init() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	envDbName := os.Getenv("DATABASE_NAME")
	envLogDbName := os.Getenv("DATABASE_LOG_NAME")
	if envDbName == "" || envLogDbName == "" {
		log.Fatal("Error on read database name.")
	}

	dbName, err := strconv.Atoi(envDbName)
	if err != nil {
		log.Fatal(err)
		panic("Cannot read envDbName")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DATABASE_ADDRESS"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DB:       dbName,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(pong)
	client = rdb
	CredentialRepo = NewRepo(ctx, rdb)
}

func CloseDb() {
	client.Close()
}
