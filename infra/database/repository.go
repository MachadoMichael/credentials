package database

import (
	"context"

	"github.com/MachadoMichael/GoAPI/schema"
	"github.com/go-redis/redis/v8"
)

type Repository interface {
	Create(credentials schema.Credentials) error
	Read(ctx context.Context, key string) (string, error)
	Delete(email string) (int64, error)
}

type Repo struct {
	ctx context.Context
	db  *redis.Client
}

func NewRepo(ctx context.Context, db *redis.Client) *Repo {
	return &Repo{ctx: ctx, db: db}
}

func (r *Repo) Create(credentials schema.Credentials) error {
	return r.db.Set(r.ctx, credentials.Email, credentials.Password, 1000000000).Err()
}

func (r *Repo) Read(key string) (string, error) {
	result, err := r.db.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key does not exist
	} else if err != nil {
		return "", err
	}
	return result, nil
}

func (r *Repo) Delete(email string) (int64, error) {
	return r.db.Del(r.ctx, email).Result()
}
