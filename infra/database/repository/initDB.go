package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var AuthRepo *BasicAuthRepo

func InitDB() {
	ctx := context.Background()

	envDbName := os.Getenv("DATABASE_NAME")

	dbName, err := strconv.Atoi(envDbName)
	if err != nil {
		log.Fatal(err)
		return
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

	AuthRepo = NewBasicAuthRepo(ctx, rdb)
}
