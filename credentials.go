package credentials

import (
	"net/http"

	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/MachadoMichael/credentials/schema"
)

type CredentialsHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	IsValidToken(w http.ResponseWriter, r *http.Request) bool
	Login(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
}

type Service struct {
	Repo         *schema.Repo
	AccessLogger *logger.Logger
	ErrorLogger  *logger.Logger
}

func NewService(repo *schema.Repo, al *logger.Logger, el *logger.Logger) CredentialsHandler {
	return &Service{
		Repo:         repo,
		AccessLogger: al,
		ErrorLogger:  el,
	}
}
