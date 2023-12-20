package user

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"github.com/ecommerce/config"
	"github.com/ecommerce/domain/auth/repository"
)

type DB struct {
	Dbx   *sqlx.DB
	Redis *redis.Client
	cfg   config.JWT
}

func ServiceUser(router fiber.Router, db DB) {
	authRepository := repository.NewAuthRepository(db.Dbx)
	redisRepository := repository.NewRedisRepository(db.Redis)

	service := NewUserService(authRepository, redisRepository)
}
