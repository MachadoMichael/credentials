package handler

import (
	"log"
	"net/http"

	"github.com/MachadoMichael/GoAPI/dto"
	"github.com/MachadoMichael/GoAPI/infra/database"
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
	if err != nil || credentialPassword != request.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Login successfully.")

}

func Create(ctx *gin.Context) {
	request := schema.Credentials{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.CredentialRepo.Create(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create credential"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Credential created successfully"})
}

func Delete(ctx *gin.Context) {
	email := ctx.Param("email")

	rows, err := database.CredentialRepo.Delete(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete credential"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Credential deleted successfully", "rows_affcteds": rows})
}

func UpdatePassword(ctx *gin.Context) {
	request := dto.UpdatePasswordRequest{}

	cred := schema.Credentials{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	credentialPassword, err := database.CredentialRepo.Read(request.Email)
	if err != nil || credentialPassword != request.OldPassword {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	cred.Email = request.Email
	cred.Password = request.OldPassword

	rows, err := database.CredentialRepo.Delete(request.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update credential", "rows_affcteds": rows})
	}

	err = database.CredentialRepo.Create(schema.Credentials{Email: request.Email, Password: request.NewPassword})
	if err != nil {
		backErr := database.CredentialRepo.Create(cred)
		if backErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create backup data credential"})
			log.Fatal("cannot save backup %i", cred)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create credential"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Password update successful"})

}