package middleware

import (
	"fmt"
	"testing"

	"github.com/ecommerce/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

var jwtSecret config.JWT

func TestMain(m *testing.M) {
	err := config.LoadConfig("../../config/config.yaml")
	if err != nil {
		panic(err)
	}
	jwtSecret = config.Cfg.JWT
	SetJWTSecretKey(jwtSecret.Secret)
	m.Run()
}

func TestGenerateNewJWT(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		ID:    "1",
		Email: "user@gmail.com",
		Role:  "user",
	})
	signedToken, err := token.SignedString([]byte(jwtSecret.Secret))

	require.NoError(t, err)
	require.NotEmpty(t, signedToken)
	require.NotEqual(t, "", signedToken)
}

func TestGetJwtClaims(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		ID:    "1",
		Email: "user@gmail.com",
		Role:  "user",
	})
	signedToken, err := token.SignedString([]byte(jwtSecret.Secret))
	require.NoError(t, err)

	ctx.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", signedToken))

	tokenHeader := ctx.Get("Authorization")

	claims, err := GetJWTClaims(tokenHeader)
	id := claims.ID
	ctx.Locals("id", id)

	require.NoError(t, err)
	require.NotEmpty(t, claims.ID)
	require.NotEqual(t, "", claims.ID)
}
