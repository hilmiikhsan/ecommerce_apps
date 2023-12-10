package auth

import (
	"github.com/ecommerce/config"
	"github.com/ecommerce/domain/auth/repository"
	"github.com/ecommerce/infra/middleware"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Dbx   *sqlx.DB
	Redis *redis.Client
	Cfg   config.JWT
}

func RegisterServiceAuth(router fiber.Router, db DB) {
	authRepository := repository.NewAuthRepository(db.Dbx)
	redisRepository := repository.NewRedisRepository(db.Redis)
	service := NewAuthService(authRepository, redisRepository, db.Cfg)
	handler := NewAuthHandler(service)

	var authRouter = router.Group("/v1/auth")
	{
		authRouter.Post("/register", handler.Register)
		authRouter.Post("/login", handler.Login)
		authRouter.Patch("/role", middleware.AuthMiddleware(), handler.UpdateRole)
	}
}
