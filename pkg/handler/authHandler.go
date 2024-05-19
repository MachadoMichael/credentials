package handler

import (
	"fmt"
	"net/http"

	"github.com/MachadoMichael/GoAPI/dto"
	"github.com/MachadoMichael/GoAPI/infra/database/repository"
	"github.com/MachadoMichael/GoAPI/schema"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	request := schema.Credentials{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creds, err := repository.AuthRepo.Login(request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, creds)

	fmt.Printf("values in request are email %s, password %s", request.Email, request.Password)
}

func Create(ctx *gin.Context) {
	request := schema.Credentials{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.AuthRepo.Create(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create credential"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Credential created successfully"})
}

func Delete(ctx *gin.Context) {
	email := ctx.Param("email")

	err := repository.AuthRepo.Delete(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete credential"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Credential deleted successfully"})
}

func UpdatePassword(ctx *gin.Context) {
	request := dto.UpdatePasswordRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.AuthRepo.UpdatePassword(request.Email, request.OldPassword, request.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password update successful"})

}
