package handler

import (
	"net/http"

	"github.com/MachadoMichael/GoAPI/infra/database"
	"github.com/MachadoMichael/GoAPI/schema"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	request := schema.Credentials{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.CredentialRepo.Create(request)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Credential created successfully"})
}
