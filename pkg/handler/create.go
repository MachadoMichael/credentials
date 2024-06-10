package handler

import (
	"net/http"
	"unicode/utf8"

	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/MachadoMichael/credentials/schema"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Create(ctx *gin.Context) {
	request := schema.Credentials{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	cred, err := database.CredentialRepo.ReadOne(request.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	if cred != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "This email is already in use"})
		logger.ErrorLogger.Write(slog.LevelError, "This email is already iin use")
		return
	}

	if utf8.RuneCountInString(request.Password) < 6 {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "The password must be longer than 6 characters."})
		logger.ErrorLogger.Write(slog.LevelError, "The password must be longer than 6 characters.")
		return
	}

	hash, hashErr := encrypt.HashPassword(request.Password)
	if hashErr != nil {
		logger.ErrorLogger.Write(slog.LevelError, hashErr.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": hashErr.Error})
		return
	}

	request.Password = hash
	errDb := database.CredentialRepo.Create(request)
	if errDb != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errDb.Error})
		logger.ErrorLogger.Write(slog.LevelError, errDb.Error())
		return
	}

	logger.AccessLogger.Write(slog.LevelInfo, "New Credential created successfully, email: "+request.Email)
	ctx.JSON(http.StatusOK, gin.H{"message": "Credential created successfully"})
}
