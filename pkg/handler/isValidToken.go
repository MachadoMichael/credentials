package handler

import (
	"net/http"
	"strings"

	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func isValidToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
		ctx.Abort()
		logger.ErrorLogger.Write(slog.LevelError, "no token provided")
		return false
	}

	strippedTokenStr := strings.TrimPrefix(token, "Bearer ")

	res, err := encrypt.ValidateToken(strippedTokenStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return false
	}

	logger.AccessLogger.Write(slog.LevelInfo, "Token validate successfully, token "+token)
	ctx.JSON(http.StatusOK, gin.H{"message": "token authorized successfully."})
	return res

}
