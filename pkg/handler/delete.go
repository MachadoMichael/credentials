package handler

import (
	"net/http"

	"github.com/MachadoMichael/GoAPI/infra/database"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	email := ctx.Param("email")

	if !isValidToken(ctx) {
		return
	}

	cred, errRead := database.CredentialRepo.Read(email)
	if errRead != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errRead.Error})
		return
	}

	if cred == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There isn't register for this email."})
		return
	}

	rows, err := database.CredentialRepo.Delete(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Credential deleted successfully", "rows_affcteds": rows})
}
