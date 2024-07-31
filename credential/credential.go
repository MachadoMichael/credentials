package credential

import (
	"net/http"

	"github.com/MachadoMichael/credentials/model"
	"github.com/MachadoMichael/credentials/pkg/logger"
)

type credentialService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	IsValidToken(w http.ResponseWriter, r *http.Request) bool
	Login(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
}

type credentialHandler struct {
	Repo         model.RepoService
	AccessLogger *logger.Logger
	ErrorLogger  *logger.Logger
}

func NewHandler(r model.RepoService, al *logger.Logger, el *logger.Logger) credentialService {
	return &credentialHandler{
		Repo:         r,
		AccessLogger: al,
		ErrorLogger:  el,
	}
}
