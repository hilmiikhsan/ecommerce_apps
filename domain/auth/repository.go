package auth

import (
	"context"

	"github.com/ecommerce/entity"
)

type Repository interface {
	Create(ctx context.Context, user entity.Auth) (err error)
	GetByEmail(ctx context.Context, email string) (user entity.Auth, err error)
	UpdateRole(ctx context.Context, id string) (err error)
}

type RedisRepository interface {
	Set(ctx context.Context, timeLimit int, token, id, email string) (err error)
	Get(ctx context.Context, id, email string) (token string, err error)
}
