package handler

import (
	"net/http"

	"github.com/MachadoMichael/GoAPI/infra/database"
	"github.com/MachadoMichael/GoAPI/pkg/encrypt"
	"github.com/MachadoMichael/GoAPI/schema"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	request := schema.Credentials{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	credentialPassword, err := database.CredentialRepo.Read(request.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = encrypt.VerifyPassword(request.Password, credentialPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := encrypt.GenerateToken(request.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successfully.", "token": token})

}
