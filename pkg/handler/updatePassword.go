package handler

import (
	"net/http"

	"github.com/MachadoMichael/credentials/dto"
	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/MachadoMichael/credentials/schema"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func UpdatePassword(ctx *gin.Context) {

	if !isValidToken(ctx) {
		return
	}

	request := dto.UpdatePasswordRequest{}
	credBackup := schema.Credentials{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	credentialPassword, err := database.CredentialRepo.Read(request.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	err = encrypt.VerifyPassword(request.OldPassword, credentialPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	credBackup.Email = request.Email
	credBackup.Password = request.OldPassword

	rows, err := database.CredentialRepo.Delete(request.Email)
	if err != nil {
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error, "rows_affcteds": rows})
	}

	err = database.CredentialRepo.Create(schema.Credentials{Email: request.Email, Password: request.NewPassword})
	if err != nil {
		backErr := database.CredentialRepo.Create(credBackup)
		if backErr != nil {
			logger.ErrorLogger.Write(slog.LevelError, backErr.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"error": backErr.Error})
		}
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}

	logger.LoginLogger.Write(slog.LevelInfo, "Successful password update, email: "+request.Email)
	ctx.JSON(http.StatusOK, gin.H{"message": "Password update successfully"})

}
