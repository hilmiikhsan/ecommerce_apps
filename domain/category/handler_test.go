package category

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/ecommerce/config"
	"github.com/ecommerce/dto"
	"github.com/ecommerce/entity"
	"github.com/ecommerce/infra/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

var handler = CategoryHandler{}

type mockCategoryService struct{}

// CreateCategory implements Service.
func (mockCategoryService) CreateCategory(ctx context.Context, req entity.Category) (err error) {
	return CreateCategoryHandler()
}

// GetListCategory implements Service.
func (mockCategoryService) GetListCategory(ctx context.Context) (response []dto.GetListCategoryResponse, err error) {
	return GetListCategoryHandler()
}

var (
	CreateCategoryHandler  func() (err error)
	GetListCategoryHandler func() (response []dto.GetListCategoryResponse, err error)
	jwtSecret              config.JWT
)

func init() {
	mock := mockCategoryService{}

	handler = NewCategoryHandler(mock)
}

func TestMain(m *testing.M) {
	err := config.LoadConfig("../../config/config.yaml")
	if err != nil {
		panic(err)
	}
	jwtSecret = config.Cfg.JWT
	middleware.SetJWTSecretKey(jwtSecret.Secret)
	m.Run()
}

func TestCreateCategoryHandler(t *testing.T) {
	type testCase struct {
		title              string
		expectedErr        error
		request            dto.CreateCategoryRequest
		expectedStatusCode int
		endpoint           string
		contentType        string
		requestHeader      string
		before             func() error
	}

	var testCases = []testCase{
		{
			title:       "create category success",
			expectedErr: nil,
			request: dto.CreateCategoryRequest{
				Name: "category 1",
			},
			expectedStatusCode: fiber.StatusCreated,
			endpoint:           "/v1/categories",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateCategoryHandler = func() (err error) {
					return nil
				}
				return nil
			},
		},
		{
			title:       "create category failed unauthorized",
			expectedErr: middleware.ErrUnAuthorized,
			request: dto.CreateCategoryRequest{
				Name: "test",
			},
			expectedStatusCode: fiber.StatusUnauthorized,
			endpoint:           "/v1/categories",
			contentType:        "application/json",
			requestHeader:      "",
			before: func() error {
				CreateCategoryHandler = func() (err error) {
					return middleware.ErrUnAuthorized
				}

				return middleware.ErrUnAuthorized
			},
		},
		{
			title:       "create category failed endpoint not found",
			expectedErr: errors.New("endpoint not found"),
			request: dto.CreateCategoryRequest{
				Name: "test",
			},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/category",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateCategoryHandler = func() (err error) {
					return errors.New("endpoint not found")
				}

				return errors.New("endpoint not found")
			},
		},
		{
			title:       "create category failed invalid content type",
			expectedErr: errors.New("invalid content type"),
			request: dto.CreateCategoryRequest{
				Name: "test",
			},
			expectedStatusCode: fiber.StatusInternalServerError,
			endpoint:           "/v1/categories",
			contentType:        "application/xml",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateCategoryHandler = func() (err error) {
					return errors.New("invalid content type")
				}

				return errors.New("invalid content type")
			},
		},
		{
			title:       "create category failed name is required",
			expectedErr: entity.ErrCategoryNameIsRequired,
			request: dto.CreateCategoryRequest{
				Name: "",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/categories",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateCategoryHandler = func() (err error) {
					return entity.ErrCategoryNameIsRequired
				}

				return entity.ErrCategoryNameIsRequired
			},
		},
		{
			title:       "create category failed internal server error",
			expectedErr: errors.New("internal server error"),
			request: dto.CreateCategoryRequest{
				Name: "test",
			},
			expectedStatusCode: fiber.StatusInternalServerError,
			endpoint:           "/v1/categories",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateCategoryHandler = func() (err error) {
					return errors.New("internal server error")
				}

				return errors.New("internal server error")
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			router := fiber.New()

			beforeErr := test.before()

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.Claims{
				ID:    "1",
				Email: "user@gmail.com",
				Role:  "user",
			})
			signedToken, err := token.SignedString([]byte(jwtSecret.Secret))
			require.NoError(t, err)

			mockService := mockCategoryService{}
			handler := NewCategoryHandler(mockService)

			router.Post("/v1/categories", middleware.AuthMiddleware(), handler.CreateCategory)

			reqBody, _ := json.Marshal(test.request)

			request := httptest.NewRequest(fiber.MethodPost, test.endpoint, bytes.NewBuffer(reqBody))
			request.Header.Set("Content-Type", test.contentType)
			request.Header.Set("Authorization", test.requestHeader+signedToken)

			resp, _ := router.Test(request, 1)

			require.Equal(t, test.expectedStatusCode, resp.StatusCode)
			require.Equal(t, test.expectedErr, beforeErr)
		})
	}
}

func TestGetListCategoryHandler(t *testing.T) {
	type testCase struct {
		title              string
		expectedErr        error
		expectedValue      []dto.GetListCategoryResponse
		expectedStatusCode int
		endpoint           string
		requestHeader      string
		before             func() error
	}

	var testCases = []testCase{
		{
			title:       "get list category success",
			expectedErr: nil,
			expectedValue: []dto.GetListCategoryResponse{
				{
					ID:   1,
					Name: "category 1",
				},
				{
					ID:   2,
					Name: "category 2",
				},
			},
			expectedStatusCode: fiber.StatusOK,
			endpoint:           "/v1/categories",
			requestHeader:      "Bearer ",
			before: func() error {
				GetListCategoryHandler = func() (response []dto.GetListCategoryResponse, err error) {
					return []dto.GetListCategoryResponse{
						{
							ID:   1,
							Name: "category 1",
						},
						{
							ID:   2,
							Name: "category 2",
						},
					}, nil
				}
				return nil
			},
		},
		{
			title:              "create category failed unauthorized",
			expectedErr:        middleware.ErrUnAuthorized,
			expectedValue:      []dto.GetListCategoryResponse{},
			expectedStatusCode: fiber.StatusUnauthorized,
			endpoint:           "/v1/categories",
			requestHeader:      "",
			before: func() error {
				GetListCategoryHandler = func() (response []dto.GetListCategoryResponse, err error) {
					return []dto.GetListCategoryResponse{}, middleware.ErrUnAuthorized
				}

				return middleware.ErrUnAuthorized
			},
		},
		{
			title:              "create category failed endpoint not found",
			expectedErr:        errors.New("endpoint not found"),
			expectedValue:      []dto.GetListCategoryResponse{},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/category",
			requestHeader:      "Bearer ",
			before: func() error {
				GetListCategoryHandler = func() (response []dto.GetListCategoryResponse, err error) {
					return []dto.GetListCategoryResponse{}, errors.New("endpoint not found")
				}

				return errors.New("endpoint not found")
			},
		},
		{
			title:              "create category failed internal server error",
			expectedErr:        errors.New("internal server error"),
			expectedValue:      []dto.GetListCategoryResponse{},
			expectedStatusCode: fiber.StatusInternalServerError,
			endpoint:           "/v1/categories",
			requestHeader:      "Bearer ",
			before: func() error {
				GetListCategoryHandler = func() (response []dto.GetListCategoryResponse, err error) {
					return []dto.GetListCategoryResponse{}, errors.New("internal server error")
				}

				return errors.New("internal server error")
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			router := fiber.New()

			beforeErr := test.before()

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.Claims{
				ID:    "1",
				Email: "user@gmail.com",
				Role:  "user",
			})
			signedToken, err := token.SignedString([]byte(jwtSecret.Secret))
			require.NoError(t, err)

			mockService := mockCategoryService{}
			handler := NewCategoryHandler(mockService)

			router.Get("/v1/categories", middleware.AuthMiddleware(), handler.GetListCategory)

			request := httptest.NewRequest(fiber.MethodGet, test.endpoint, nil)
			request.Header.Set("Authorization", test.requestHeader+signedToken)

			resp, _ := router.Test(request, 1)

			require.Equal(t, test.expectedStatusCode, resp.StatusCode)
			require.Equal(t, test.expectedErr, beforeErr)
		})
	}
}
