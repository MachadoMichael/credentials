package main

import (
	"github.com/MachadoMichael/credentials/infra"
	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/MachadoMichael/credentials/pkg/route"
)

func main() {
	infra.Init()
	logger.InitLoggers()
	defer logger.AccessLogger.Close()
	defer logger.ErrorLogger.Close()

	database.Init()
	defer database.CloseDb()
	route.Init()
}
