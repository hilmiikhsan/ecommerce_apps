package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/ecommerce/config"
	"github.com/ecommerce/dto"
	"github.com/ecommerce/entity"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

var svc = AuthService{}

type mockAuthRepository struct{}
type mockRedisRepository struct{}

// Get implements RedisRepository.
func (mockRedisRepository) Get(ctx context.Context, id string, email string) (token string, err error) {
	return Get()
}

// Set implements RedisRepository.
func (mockRedisRepository) Set(ctx context.Context, timeLimit int, token string, id string, email string) (err error) {
	return Set()
}

// Create implements Repository.
func (mockAuthRepository) Create(ctx context.Context, user entity.Auth) (err error) {
	return Create()
}

// GetByEmail implements Repository.
func (mockAuthRepository) GetByEmail(ctx context.Context, email string) (user entity.Auth, err error) {
	return GetByEmail()
}

// UpdateRole implements Repository.
func (mockAuthRepository) UpdateRole(ctx context.Context, id string) (err error) {
	return nil
}

func EncryptPassword(password string) (result string, err error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	password = string(encrypted)
	return password, nil
}

var (
	Get        func() (token string, err error)
	Set        func() (err error)
	Create     func() (err error)
	GetByEmail func() (user entity.Auth, err error)
	UpdateRole func() (err error)
)

func init() {
	mock := mockAuthRepository{}
	mockRedis := mockRedisRepository{}
	jwtConfig := config.JWT{
		Secret:            "secret",
		TokenLifeTimeHour: 7,
	}

	svc = NewAuthService(mock, mockRedis, jwtConfig)
}

func TestRegister(t *testing.T) {
	type testCase struct {
		title       string
		expectedErr error
		request     entity.Auth
		before      func()
	}

	var testCases = []testCase{
		{
			title:       "register success",
			expectedErr: nil,
			request: entity.Auth{
				ID:       "1",
				Email:    "user2@gmail.com",
				Password: "password",
			},
			before: func() {
				GetByEmail = func() (user entity.Auth, err error) {
					return entity.Auth{
						ID:       "1",
						Email:    "user@gmail.com",
						Password: "password",
					}, nil
				}

				Create = func() (err error) {
					return nil
				}
			},
		},
		{
			title:       "register failed email already used",
			expectedErr: errors.New("email already used"),
			request: entity.Auth{
				ID:       "1",
				Email:    "user@gmail.com",
				Password: "password",
			},
			before: func() {
				GetByEmail = func() (user entity.Auth, err error) {
					return entity.Auth{
						ID:       "1",
						Email:    "user@gmail.com",
						Password: "password",
					}, nil
				}

				Create = func() (err error) {
					return errors.New("email already used")
				}
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.before()

			err := svc.Register(context.Background(), test.request)
			require.Equal(t, test.expectedErr, err)
		})
	}
}

func TestLogin(t *testing.T) {
	type testCase struct {
		title         string
		expectedErr   error
		expectedValue dto.LoginResponse
		request       entity.Auth
		before        func()
	}

	var testCases = []testCase{
		{
			title:       "login success",
			expectedErr: nil,
			expectedValue: dto.LoginResponse{
				AccessToken: "token",
				Role:        "user",
			},
			request: entity.Auth{
				Email:    "user@gmail.com",
				Password: "password",
			},
			before: func() {
				password, _ := EncryptPassword("password")

				GetByEmail = func() (user entity.Auth, err error) {
					return entity.Auth{
						ID:       "1",
						Email:    "user@gmail.com",
						Role:     "user",
						Password: password,
					}, nil
				}

				Get = func() (token string, err error) {
					return "token", nil
				}

				Set = func() (err error) {
					return nil
				}
			},
		},
		{
			title:         "login failed invalid email or password",
			expectedErr:   entity.ErrInvalidEmailOrPassword,
			expectedValue: dto.LoginResponse{},
			request: entity.Auth{
				Email:    "user2@gmail.com",
				Password: "password",
			},
			before: func() {
				password, _ := EncryptPassword("password2")

				GetByEmail = func() (user entity.Auth, err error) {
					return entity.Auth{
						ID:       "1",
						Email:    "user@gmail.com",
						Role:     "user",
						Password: password,
					}, nil
				}

				Get = func() (token string, err error) {
					return "", nil
				}

				Set = func() (err error) {
					return nil
				}
			},
		},
		{
			title:         "login failed: get access token error",
			expectedValue: dto.LoginResponse{},
			expectedErr:   errors.New("get access token error"),
			request: entity.Auth{
				Email:    "user@gmail.com",
				Password: "password",
			},
			before: func() {
				password, _ := EncryptPassword("password")

				GetByEmail = func() (user entity.Auth, err error) {
					return entity.Auth{
						ID:       "1",
						Email:    "user@gmail.com",
						Role:     "user",
						Password: password,
					}, nil
				}

				Get = func() (token string, err error) {
					return "", errors.New("get access token error")
				}

				Set = func() (err error) {
					return nil
				}
			},
		},
		{
			title:         "login failed: internal server error",
			expectedValue: dto.LoginResponse{},
			expectedErr:   errors.New("internal server error"),
			request: entity.Auth{
				Email:    "user2@gmail.com",
				Password: "password",
			},
			before: func() {
				GetByEmail = func() (user entity.Auth, err error) {
					return entity.Auth{}, errors.New("internal server error")
				}

				Get = func() (token string, err error) {
					return "", nil
				}

				Set = func() (err error) {
					return nil
				}
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.before()

			user, accessToken, err := svc.Login(context.Background(), test.request)
			require.Equal(t, test.expectedErr, err)
			require.Equal(t, test.expectedValue.AccessToken, accessToken)
			require.Equal(t, test.expectedValue.Role, user.Role)
		})
	}
}

func TestUpdateRole(t *testing.T) {
	type testCase struct {
		title       string
		expectedErr error
		id          string
		email       string
		before      func()
	}

	var testCases = []testCase{
		{
			title:       "update role success",
			expectedErr: nil,
			id:          "1",
			email:       "user@gmail.com",
			before: func() {
				GetByEmail = func() (user entity.Auth, err error) {
					return entity.Auth{
						ID:    "1",
						Email: "user@gmail.com",
						Role:  "user",
					}, nil
				}
			},
		},
		{

			title:       "update role failed: invalid role",
			expectedErr: entity.ErrUserAlreadyMerchant,
			id:          "1",
			email:       "user@gmail.com",
			before: func() {
				GetByEmail = func() (user entity.Auth, err error) {
					return entity.Auth{
						ID:    "1",
						Email: "user@gmail.com",
						Role:  "merchant",
					}, nil
				}
			},
		},
		{
			title:       "update role failed: internal server error",
			expectedErr: errors.New("internal server error"),
			id:          "1",
			email:       "user@gmail.com",
			before: func() {
				GetByEmail = func() (user entity.Auth, err error) {
					return entity.Auth{}, errors.New("internal server error")
				}
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.before()

			err := svc.UpdateRole(context.Background(), test.email)
			require.Equal(t, test.expectedErr, err)
		})
	}
}
