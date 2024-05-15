package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Credentials struct {
	Email    string `json: email`
	Password string `json password`
}

func Login(ctx *gin.Context) {
	request := Credentials{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("values in request are email %s, password %s", request.Email, request.Password)
}
