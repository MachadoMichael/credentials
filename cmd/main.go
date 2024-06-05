package main

import (
	"log"

	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/MachadoMichael/credentials/pkg/route"
)

func main() {
	err := logger.InitLoggers()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.LoginLogger.Close()
	defer logger.ErrorLogger.Close()

	database.Init()
	defer database.CloseDb()
	route.Init()
}
