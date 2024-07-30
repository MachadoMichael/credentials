package handler

import (
	"net/http"

	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/logger"
)

type CredentialsService interface {
	Create(w http.ResponseWriter, r *http.Request) error
	Delete(w http.ResponseWriter, r *http.Request) error
	Update(w http.ResponseWriter, r *http.Request) error
	IsValidToken(w http.ResponseWriter, r *http.Request) bool
	Login(w http.ResponseWriter, r *http.Request) error
	Read(w http.ResponseWriter, r *http.Request) error
}

type Service struct {
	Repo         *database.Repo
	AccessLogger *logger.Logger
	ErrorLogger  *logger.Logger
}

func NewService(repo *database.Repo) CredentialsService {
	return &Service{
		Repo: repo,
	}
}
