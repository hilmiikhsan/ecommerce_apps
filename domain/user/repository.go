package user

import (
	"context"

	"github.com/ecommerce/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetUserById(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, id string, user entity.User) error
}

type RedisRepository interface {
	Set(ctx context.Context, timeLimit int, token, id, email string) (err error)
	Get(ctx context.Context, id, email string) (token string, err error)
}
