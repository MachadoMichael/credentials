package route

import (
	"net/http"

	"github.com/MachadoMichael/credentials/pkg/handler"
)

func Init() {
	http.HandleFunc("/read", handler.Read)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/{email}", handler.Delete)
	http.HandleFunc("/update", handler.Update)
	http.ListenAndServe(":8080", nil)
}
