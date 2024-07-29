package credentials

import (
	"github.com/MachadoMichael/credentials/handler"
	"github.com/MachadoMichael/credentials/infra"
	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/logger"
)

type Credentials interface {
	Init() *handler.Service
}

type Cred struct {
}

func NewCredentialHandler() Credentials {
	return &Cred{}
}

func (c *Cred) Init() *handler.Service {
	infra.Init()
	al, el, err := logger.Init()
	if err != nil {
		panic(err)
	}

	defer al.Close()
	defer el.Close()

	database.Init()
	defer database.CloseDb()

	return &handler.Service{
		Repo:         database.CredentialRepo,
		AccessLogger: *logger.Logger,
		ErrorLogger:  *logger.Logger,
	}
}
