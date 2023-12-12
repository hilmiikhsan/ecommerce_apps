package entity

import (
	"errors"
	"regexp"

	"github.com/ecommerce/dto"
	"github.com/ecommerce/infra/middleware"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailIsRequired        = errors.New("email is required")
	ErrEmailIsInvalid         = errors.New("email is invalid")
	ErrPasswordIsEmpty        = errors.New("password is empty")
	ErrPasswordLength         = errors.New("password length must be greater than equal 6")
	ErrEmailAlreadyUsed       = errors.New("email already used")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrUserAlreadyMerchant    = errors.New("user already as a merchant")
)

var EmailPattern string = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

type Auth struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) Validate(req dto.AuthRequest) (Auth, error) {
	if req.Email == "" {
		return a, ErrEmailIsRequired
	}

	emailRegex := regexp.MustCompile(EmailPattern)
	if !emailRegex.MatchString(req.Email) {
		return a, ErrEmailIsInvalid
	}

	if req.Password == "" {
		return a, ErrPasswordIsEmpty
	}

	if len(req.Password) < 6 {
		return a, ErrPasswordLength
	}

	a.Email = req.Email
	a.Password = req.Password

	return a, nil
}

func (a Auth) CheckRequestEmail(reqEmail, existsEmail string) (err error) {
	if reqEmail == existsEmail {
		return ErrEmailAlreadyUsed
	}

	return
}

func (a Auth) CheckRegisteredEmail(reqEmail, existsEmail string) (err error) {
	if reqEmail != existsEmail {
		return ErrInvalidEmailOrPassword
	}

	return
}

func (a *Auth) EncryptPassword() (err error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	a.Password = string(encrypted)
	return
}

func (a Auth) ValidatePasswordFromPlainText(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a Auth) GenerateAccessToken(accessToken string) (token string, err error) {
	token, err = middleware.GenerateNewJWT(&middleware.Claims{
		ID:    a.ID,
		Email: a.Email,
		Role:  a.Role,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a Auth) ValidateUserRole(role string) (err error) {
	if role == "merchant" {
		return ErrUserAlreadyMerchant
	}

	return
}
