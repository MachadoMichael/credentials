package route

import (
	"net/http"

	"github.com/MachadoMichael/credentials/pkg/handler"
	"github.com/gin-gonic/gin"
)

func Init() {

	http.HandleFunc("GET /read", handler.Read)
	http.HandleFunc("POST /login", handler.Login)
	http.HandleFunc("POST /create", handler.Create)
	http.HandleFunc("DELETE /{email}", handler.Delete)
	http.HandleFunc("PUT /update ", handler.Update)
	http.ListenAndServe(":8080", nil)
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/login", handler.Login)
		v1.POST("/create", handler.Create)
		v1.DELETE("/:email", handler.Delete)
		v1.PUT("/update-password", handler.UpdatePassword)
	}

	router.Run()

}
