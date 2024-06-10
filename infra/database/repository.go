package database

import (
	"context"

	"github.com/MachadoMichael/credentials/schema"
	"github.com/go-redis/redis/v8"
)

type Repo struct {
	ctx context.Context
	db  *redis.Client
}

func NewRepo(ctx context.Context, db *redis.Client) *Repo {
	return &Repo{ctx: ctx, db: db}
}

func (r *Repo) Create(credentials schema.Credentials) error {
	return r.db.Set(r.ctx, credentials.Email, credentials.Password, 0).Err()
}

func (r *Repo) Read() (map[string]string, error) {
	keys := make(map[string]string)
	var cursor uint64

	for {
		var err error
		var keysPart []string

		keysPart, cursor, err = r.db.Scan(r.ctx, cursor, "*", 10).Result()
		if err != nil {
			return nil, err
		}

		for _, key := range keysPart {
			value, err := r.db.Get(r.ctx, key).Result()
			if err != nil {
				if err == redis.Nil {
					keys[key] = ""
				} else {
					return nil, err
				}
			} else {
				keys[key] = value
			}
		}
		if cursor == 0 {
			break
		}

	}
	return keys, nil
}

func (r *Repo) ReadOne(key string) (string, error) {
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
