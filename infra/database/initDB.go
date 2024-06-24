package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MachadoMichael/credentials/infra"
	"github.com/go-redis/redis/v8"
)

var CredentialRepo *Repo
var client *redis.Client

func Init() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DATABASE_ADDRESS"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DB:       infra.Config.DbName,
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
