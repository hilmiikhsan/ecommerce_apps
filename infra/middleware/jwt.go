package middleware

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const (
	bearer = "Bearer"
)

var (
	ErrUnAuthorized = errors.New("unauthorized")
)

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecretKey = ""

func SetJWTSecretKey(key string) {
	jwtSecretKey = key
}

func GenerateNewJWT(claims *Claims) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	signedToken, err = token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GetJWTClaims(tokenString string) (claims *Claims, err error) {
	splitToken := strings.Split(tokenString, bearer)
	if len(splitToken) != 2 {
		return nil, ErrUnAuthorized
	}
	reqToken := strings.TrimSpace(splitToken[1])

	claims = &Claims{}
	token, err := jwt.ParseWithClaims(reqToken, claims, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, ErrUnAuthorized
	}

	if !token.Valid {
		return nil, ErrUnAuthorized
	}
	return claims, nil
}
