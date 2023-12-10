package middleware

import (
	"fmt"

	logs "github.com/ecommerce/infra/logger"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenHeader := ctx.Get("Authorization")
		if tokenHeader == "" {
			logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", ErrUnAuthorized))
			return WriteError(ctx, ErrUnAuthorized)
		}

		claims, err := GetJWTClaims(tokenHeader)
		if err != nil {
			logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
			return WriteError(ctx, err)
		}

		id := claims.ID
		email := claims.Email

		ctx.Locals("id", id)
		ctx.Locals("email", email)

		return ctx.Next()
	}
}
