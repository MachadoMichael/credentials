package repository

import (
	"context"

	"github.com/MachadoMichael/GoAPI/schema"
	"github.com/go-redis/redis/v8"
)

type AuthRepository interface {
	Create(credentials schema.Credentials) error
	Read(ctx context.Context, key string) (string, error)
	Delete(email string) (int64, error)
}

type BasicAuthRepo struct {
	ctx context.Context
	db  *redis.Client
}

func NewBasicAuthRepo(ctx context.Context, db *redis.Client) *BasicAuthRepo {
	return &BasicAuthRepo{ctx: ctx, db: db}
}

func (b *BasicAuthRepo) Create(credentials schema.Credentials) error {
	return b.db.Set(b.ctx, credentials.Email, credentials.Password, 0).Err()
}

func (b *BasicAuthRepo) Read(ctx context.Context, key string) (string, error) {
	result, err := b.db.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key does not exist
	} else if err != nil {
		return "", err
	}
	return result, nil
}

func (b *BasicAuthRepo) Delete(email string) (int64, error) {
	return b.db.Del(b.ctx, email).Result()
}
