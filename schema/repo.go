package schema

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Repo struct {
	ctx context.Context
	db  *redis.Client
}
