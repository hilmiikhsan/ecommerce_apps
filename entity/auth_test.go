package entity

import (
	"testing"

	"github.com/ecommerce/dto"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestEntityAuth(t *testing.T) {
	t.Run("err : email is required", func(t *testing.T) {
		var req = dto.AuthRequest{
			Email:    "",
			Password: "123456",
		}

		_, err := NewAuth().Validate(req)
		require.NotNil(t, err)
		require.Equal(t, ErrEmailIsRequired, err)
	})

	t.Run("err : email is invalid", func(t *testing.T) {
		var req = dto.AuthRequest{
			Email:    "test",
			Password: "123456",
		}

		_, err := NewAuth().Validate(req)
		require.NotNil(t, err)
		require.Equal(t, ErrEmailIsInvalid, err)
	})

	t.Run("err : password is empty", func(t *testing.T) {
		var req = dto.AuthRequest{
			Email:    "test@gmail.com",
			Password: "",
		}

		_, err := NewAuth().Validate(req)
		require.NotNil(t, err)
		require.Equal(t, ErrPasswordIsEmpty, err)
	})

	t.Run("err : password length must be greater than equal 6", func(t *testing.T) {
		var req = dto.AuthRequest{
			Email:    "test@gmail.com",
			Password: "123",
		}

		_, err := NewAuth().Validate(req)
		require.NotNil(t, err)
		require.Equal(t, ErrPasswordLength, err)
	})

	t.Run("success : validate auth request", func(t *testing.T) {
		var req = dto.AuthRequest{
			Email:    "test@gmail.com",
			Password: "123456",
		}

		_, err := NewAuth().Validate(req)
		require.Nil(t, err)
		require.Equal(t, nil, err)
	})

	t.Run("err : check request email", func(t *testing.T) {
		var reqEmail = "test@gmail.com"
		var existsEmail = "test@gmail.com"

		err := NewAuth().CheckRequestEmail(reqEmail, existsEmail)
		require.NotNil(t, err)
		require.Equal(t, ErrEmailAlreadyUsed, err)
	})

	t.Run("err : check registered email", func(t *testing.T) {
		var reqEmail = "test2@gmail.com"
		var existsEmail = "test@gmail.com"

		err := NewAuth().CheckRegisteredEmail(reqEmail, existsEmail)
		require.NotNil(t, err)
		require.Equal(t, ErrInvalidEmailOrPassword, err)
	})

	t.Run("success : check request email", func(t *testing.T) {
		var reqEmail = "test2@gmail.com"
		var existsEmail = "test@gmail.com"

		err := NewAuth().CheckRequestEmail(reqEmail, existsEmail)
		require.Nil(t, err)
		require.Equal(t, nil, err)
	})

	t.Run("success : check registered email", func(t *testing.T) {
		var reqEmail = "test@gmail.com"
		var existsEmail = "test@gmail.com"

		err := NewAuth().CheckRegisteredEmail(reqEmail, existsEmail)
		require.Nil(t, err)
		require.Equal(t, nil, err)
	})

	t.Run("err : encrypt password", func(t *testing.T) {
		encrypted, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		require.Nil(t, err)
		require.Equal(t, nil, err)
		require.NotEqual(t, "123456", string(encrypted))
	})

	t.Run("success : encrypt password", func(t *testing.T) {
		encrypted, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		require.Nil(t, err)
		require.Equal(t, nil, err)
		require.NotEqual(t, "123456", string(encrypted))
	})

	t.Run("err : validate password hash", func(t *testing.T) {
		password := "123456"
		hash := "$2a$10$Qz7V5xkX6q4h7xYj1QJQ4uZ5Xn0p5Zb9kK5aJwZ1jZb7n6j3z7c3S"

		err := NewAuth().ValidatePasswordFromPlainText(password, hash)
		require.NotNil(t, err)
		require.Equal(t, false, err)
	})

	t.Run("success : validate password hash", func(t *testing.T) {
		encrypted, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		password := string(encrypted)

		err := bcrypt.CompareHashAndPassword([]byte(password), []byte("123456"))
		require.Nil(t, err)

		result := NewAuth().ValidatePasswordFromPlainText("123456", password)
		require.True(t, result)
	})

	t.Run("success : generate access token", func(t *testing.T) {
		token, err := NewAuth().GenerateAccessToken("token")
		require.Nil(t, err)
		require.Equal(t, nil, err)
		require.NotEqual(t, "token", token)
	})

	t.Run("err : validate user role", func(t *testing.T) {
		err := NewAuth().ValidateUserRole("merchant")
		require.NotNil(t, err)
		require.Equal(t, ErrUserAlreadyMerchant, err)
	})

	t.Run("success : validate user role", func(t *testing.T) {
		err := NewAuth().ValidateUserRole("customer")
		require.Nil(t, err)
		require.Equal(t, nil, err)
	})
}
