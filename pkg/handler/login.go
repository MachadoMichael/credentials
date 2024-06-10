package handler

import (
	"net/http"

	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/MachadoMichael/credentials/schema"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Login(ctx *gin.Context) {
	request := schema.Credentials{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	credentialPassword, err := database.CredentialRepo.ReadOne(request.Email)
	if credentialPassword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Credential not found."})
		logger.ErrorLogger.Write(slog.LevelError, "Credential not found")
		return
	}
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	err = encrypt.VerifyPassword(credentialPassword, request.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	token, err := encrypt.GenerateToken(request.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	logger.AccessLogger.Write(slog.LevelInfo, "sucessful login attempt, email: "+request.Email)
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successfully.", "token": token})

}
