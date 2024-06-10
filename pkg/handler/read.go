package handler

import (
	"net/http"

	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Read(ctx *gin.Context) {
	if !isValidToken(ctx) {
		return
	}

	credentials, err := database.CredentialRepo.Read()
	if err != nil {
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	logger.AccessLogger.Write(slog.LevelInfo, "Successful to read credentials on database.")
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful to read credentials on database", "credentials": credentials})
}
