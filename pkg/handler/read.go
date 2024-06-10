package handler

import (
	"net/http"
	"time"

	"github.com/MachadoMichael/credentials/dto"
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

	var content []string

	for key := range credentials {
		content = append(content, key)
	}

	layout := "2006-01-02 15:04:05"
	formattedTime := time.Now().Format(layout)

	response := dto.CredentialsRegisteredDTO{
		Lenght:  len(content),
		Content: content,
		ReadAt:  formattedTime,
	}

	logger.AccessLogger.Write(slog.LevelInfo, "Successful to read credentials on database.")
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful to read credentials on database", "credentials": response})
}
