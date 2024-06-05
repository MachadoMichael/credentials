package handler

import (
	"net/http"
	"unicode/utf8"

	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/MachadoMichael/credentials/schema"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	request := schema.Credentials{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cred, errRead := database.CredentialRepo.Read(request.Email)
	if errRead != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errRead.Error})
		return
	}

	if cred != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "This email already on use"})
		return
	}

	if utf8.RuneCountInString(request.Password) < 6 {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "Password need to have more than 6 characters."})
		return
	}

	hash, hashErr := encrypt.HashPassword(request.Password)
	if hashErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": hashErr.Error})
		return
	}

	request.Password = hash

	err := database.CredentialRepo.Create(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	logger.LoginLogger.Write("create", "created new login")
	ctx.JSON(http.StatusOK, gin.H{"message": "Credential created successfully"})
}
