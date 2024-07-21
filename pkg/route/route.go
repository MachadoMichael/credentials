package route

import (
	"net/http"
	"time"

	"github.com/MachadoMichael/credentials/pkg/handler"
)

func Init() {

	mux := http.NewServeMux()

	mux.HandleFunc("/read", handler.Read)
	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/create", handler.Create)
	mux.HandleFunc("/{email}", handler.Delete)
	mux.HandleFunc("/update", handler.Update)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	srv.ListenAndServe()
}
