package handler

import (
	"log"
	"net/http"

	"github.com/MachadoMichael/credentials/dto"
	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/schema"
	"github.com/gin-gonic/gin"
)

func UpdatePassword(ctx *gin.Context) {

	if !isValidToken(ctx) {
		return
	}

	request := dto.UpdatePasswordRequest{}
	credBackup := schema.Credentials{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	credentialPassword, err := database.CredentialRepo.Read(request.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = encrypt.VerifyPassword(request.OldPassword, credentialPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	credBackup.Email = request.Email
	credBackup.Password = request.OldPassword

	rows, err := database.CredentialRepo.Delete(request.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error, "rows_affcteds": rows})
	}

	err = database.CredentialRepo.Create(schema.Credentials{Email: request.Email, Password: request.NewPassword})
	if err != nil {
		backErr := database.CredentialRepo.Create(credBackup)
		if backErr != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": backErr.Error})
			log.Fatal("cannot save backup %i", credBackup)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Password update successful"})

}
