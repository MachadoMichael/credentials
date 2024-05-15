package route

import (
	"github.com/MachadoMichael/GoAPI/pkg/handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("", handler.Login)
	}

	router.Run()

}
