package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var CredentialRepo *Repo

func Init() {
	ctx := context.Background()

	envDbName := os.Getenv("DATABASE_NAME")

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

	defer rdb.Close()

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(pong)
	CredentialRepo = NewRepo(ctx, rdb)

}
