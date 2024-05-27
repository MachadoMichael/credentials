package route

import (
	"github.com/MachadoMichael/GoAPI/pkg/handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/login", handler.Login)
		v1.POST("/create", handler.Create)
		v1.POST("/token", handler.TokenValidation)
		v1.DELETE("/:email", handler.Delete)
		v1.PUT("/update-password", handler.UpdatePassword)
	}

	router.Run()

}
