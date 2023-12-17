package user

import (
	"context"

	"github.com/ecommerce/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, user entity.User) error
	GetUserById(ctx context.Context)
	UpdateUser()
}

type RedisRepository interface {
	Set()
	Get()
}
