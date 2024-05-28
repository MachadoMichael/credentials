package handler

import (
	"net/http"
	"os"

	"github.com/MachadoMichael/GoAPI/pkg/encrypt"
	"github.com/gin-gonic/gin"
)

func TokenValidation(ctx *gin.Context) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "secret key not set"})
		ctx.Abort()
		return
	}

	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
		ctx.Abort()
		return
	}

	_, err := encrypt.ValidateToken(token, secret)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "token authorized successfully."})

}
