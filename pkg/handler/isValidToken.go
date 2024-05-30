package handler

import (
	"net/http"
	"strings"

	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/gin-gonic/gin"
)

func isValidToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
		ctx.Abort()
		return false
	}

	strippedTokenStr := strings.TrimPrefix(token, "Bearer ")

	res, err := encrypt.ValidateToken(strippedTokenStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		return false
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "token authorized successfully."})
	return res

}
