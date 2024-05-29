package handler

import (
	"net/http"
	"os"

	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/gin-gonic/gin"
)

func isValidToken(ctx *gin.Context) bool {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "secret key not set"})
		ctx.Abort()
		return false
	}

	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
		ctx.Abort()
		return false
	}

	_, err := encrypt.ValidateToken(token, secret)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		return false
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "token authorized successfully."})
	return true

}
