package handler

import (
	"net/http"

	"github.com/MachadoMichael/GoAPI/infra/database"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	email := ctx.Param("email")

	rows, err := database.CredentialRepo.Delete(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Credential deleted successfully", "rows_affcteds": rows})
}
