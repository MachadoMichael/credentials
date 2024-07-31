package handler

import (
	"net/http"

	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/MachadoMichael/credentials/schema"
)

type Service interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	IsValidToken(w http.ResponseWriter, r *http.Request) bool
	Login(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
}

type CredentialHandler struct {
	Repo         schema.RepoInterface
	AccessLogger *logger.Logger
	ErrorLogger  *logger.Logger
}

func New(r schema.RepoInterface, al *logger.Logger, el *logger.Logger) Service {
	return &CredentialHandler{
		Repo:         r,
		AccessLogger: al,
		ErrorLogger:  el,
	}
}
