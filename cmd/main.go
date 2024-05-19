package main

import (
	"log"

	"github.com/MachadoMichael/GoAPI/infra/database/repository"
	"github.com/MachadoMichael/GoAPI/pkg/route"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	repository.InitDB()
	route.Init()
}
