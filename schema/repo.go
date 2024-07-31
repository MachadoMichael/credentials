package schema

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Repo struct {
	ctx context.Context
	db  *redis.Client
}

type RepoInterface interface {
	Read() ([]Credentials, error)
	ReadOne(email string) (string, error)
	Delete(email string) (int, error)
	Create(cred Credentials) error
}

func (r *Repo) Read() ([]Credentials, error) {
	creds := []Credentials{}
	return creds, nil
}

func (r *Repo) ReadOne(email string) (string, error) {
	return "", nil
}

func (r *Repo) Delete(email string) (int, error) {
	return 0, nil
}

func (r *Repo) Create(cred Credentials) error {
	return nil
}

func NewRepo() RepoInterface {
	return &Repo{}
}
