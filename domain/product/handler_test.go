package product

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

var handler = ProductHandler{}

type mockProductService struct{}

// CreateProduct implements Service.
func (mockProductService) CreateProduct(ctx context.Context, req entity.Product, token string) (err error) {
	return CreateProductHandler()
}

// GetDetailProduct implements Service.
func (mockProductService) GetDetailProduct(ctx context.Context, id int, token string) (response dto.GetDetailProductResponse, err error) {
	return GetDetailProductHandler()
}

// GetDetailProductUserPerspective implements Service.
func (mockProductService) GetDetailProductUserPerspective(ctx context.Context, sku string) (response dto.GetDetailProductUserPerspectiveResponse, err error) {
	return GetDetailProductUserPerspectiveHandler()
}

// GetListProduct implements Service.
func (mockProductService) GetListProduct(ctx context.Context, token string, queryParam string, limit int, page int) (response []dto.GetListProductResponse, totalData int, err error) {
	return GetListProductHandler()
}

// UpdateProduct implements Service.
func (mockProductService) UpdateProduct(ctx context.Context, req entity.Product, token string) (err error) {
	return UpdateProductHandler()
}

var (
	CreateProductHandler                   func() (err error)
	GetDetailProductHandler                func() (response dto.GetDetailProductResponse, err error)
	GetDetailProductUserPerspectiveHandler func() (response dto.GetDetailProductUserPerspectiveResponse, err error)
	GetListProductHandler                  func() (response []dto.GetListProductResponse, totalData int, err error)
	UpdateProductHandler                   func() (err error)
	jwtSecret                              config.JWT
)

func init() {
	mock := mockProductService{}

	handler = NewProductHandler(mock)
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

// func TestCreateProductHandler_Success(t *testing.T) {
// 	router := fiber.New()

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.Claims{
// 		ID:    "1",
// 		Email: "user@gmail.com",
// 		Role:  "user",
// 	})
// 	signedToken, err := token.SignedString([]byte(jwtSecret.Secret))
// 	require.NoError(t, err)

// 	req := dto.CreateOrUpdateProductRequest{
// 		ID:          1,
// 		Name:        "test",
// 		Description: "test",
// 		Price:       1000,
// 		Stock:       100,
// 		CategoryId:  3,
// 		ImageUrl:    "test.png",
// 	}

// 	mockService := mockProductService{}

// 	handler := NewProductHandler(mockService)

// 	router.Post("/v1/products", middleware.AuthMiddleware(), handler.CreateProducts)

// 	reqBody, _ := json.Marshal(req)

// 	request := httptest.NewRequest("POST", "/v1/products", bytes.NewBuffer(reqBody))

// 	request.Header.Set("Content-Type", "application/json")
// 	request.Header.Set("Authorization", "Bearer "+signedToken)

// 	resp, _ := router.Test(request, 1)

// 	require.Equal(t, 201, resp.StatusCode)
// }

func TestCreateProductHandler(t *testing.T) {
	type testCase struct {
		title              string
		expectedErr        error
		request            dto.CreateOrUpdateProductRequest
		expectedStatusCode int
		endpoint           string
		contentType        string
		requestHeader      string
		before             func() error
	}

	var testCases = []testCase{
		{
			title:       "create product success",
			expectedErr: nil,
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusCreated,
			endpoint:           "/v1/products",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return nil
				}

				return nil
			},
		},
		{
			title:       "create product failed unauthorized",
			expectedErr: middleware.ErrUnAuthorized,
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusUnauthorized,
			endpoint:           "/v1/products",
			contentType:        "application/json",
			requestHeader:      "",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return middleware.ErrUnAuthorized
				}

				return middleware.ErrUnAuthorized
			},
		},
		{
			title:       "create product failed endpoint not found",
			expectedErr: errors.New("endpoint not found"),
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/product",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return errors.New("endpoint not found")
				}

				return errors.New("endpoint not found")
			},
		},
		{
			title:       "create product failed invalid content type",
			expectedErr: errors.New("invalid content type"),
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/product",
			contentType:        "application/xml",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return errors.New("invalid content type")
				}

				return errors.New("invalid content type")
			},
		},
		{
			title:       "create product failed name is required",
			expectedErr: entity.ErrProductNameIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "",
				Description: "test",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return entity.ErrProductNameIsRequired
				}

				return entity.ErrProductNameIsRequired
			},
		},
		{
			title:       "create product failed description is required",
			expectedErr: entity.ErrDescriptionIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "test",
				Description: "",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return entity.ErrDescriptionIsRequired
				}

				return entity.ErrDescriptionIsRequired
			},
		},
		{
			title:       "create product failed price is required",
			expectedErr: entity.ErrPriceIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "test",
				Description: "test",
				Price:       0,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return entity.ErrPriceIsRequired
				}

				return entity.ErrPriceIsRequired
			},
		},
		{
			title:       "create product failed price is invalid",
			expectedErr: entity.ErrPriceIsInvalid,
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "test",
				Description: "test",
				Price:       -1,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return entity.ErrPriceIsInvalid
				}

				return entity.ErrPriceIsInvalid
			},
		},
		{
			title:       "create product failed stock is required",
			expectedErr: entity.ErrStockIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       0,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return entity.ErrStockIsRequired
				}

				return entity.ErrStockIsRequired
			},
		},
		{
			title:       "create product failed stock is invalid",
			expectedErr: entity.ErrStockIsInvalid,
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       -1,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return entity.ErrStockIsInvalid
				}

				return entity.ErrStockIsInvalid
			},
		},
		{
			title:       "create product failed category id is required",
			expectedErr: entity.ErrCategoryIdIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       10,
				CategoryId:  0,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return entity.ErrCategoryIdIsRequired
				}

				return entity.ErrCategoryIdIsRequired
			},
		},
		{
			title:       "create product failed image url is required",
			expectedErr: entity.ErrImageUrlIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				ID:          1,
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       10,
				CategoryId:  1,
				ImageUrl:    "",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
					return entity.ErrImageUrlIsRequired
				}

				return entity.ErrImageUrlIsRequired
			},
		},
		{
			title:       "create product failed internal server error",
			expectedErr: errors.New("internal server error"),
			request: dto.CreateOrUpdateProductRequest{
				ID:          -999999999999,
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       10,
				CategoryId:  1,
				ImageUrl:    "image.png",
			},
			expectedStatusCode: fiber.StatusInternalServerError,
			endpoint:           "/v1/products",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				CreateProductHandler = func() (err error) {
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

			mockService := mockProductService{}
			handler := NewProductHandler(mockService)

			router.Post("/v1/products", middleware.AuthMiddleware(), handler.CreateProducts)

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

func TestGetDetailProductHandler(t *testing.T) {
	type testCase struct {
		title              string
		expectedErr        error
		expectedValue      dto.GetDetailProductResponse
		expectedStatusCode int
		endpoint           string
		requestHeader      string
		before             func() error
	}

	var testCases = []testCase{
		{
			title:       "get detail product success",
			expectedErr: nil,
			expectedValue: dto.GetDetailProductResponse{
				ID:          1,
				Sku:         "test",
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       10,
				Category:    "test",
				CategoryId:  1,
				ImageUrl:    "test.png",
				CreatedAt:   "2021-01-01T00:00:00Z",
				UpdatedAt:   "2021-01-01T00:00:00Z",
			},
			expectedStatusCode: fiber.StatusOK,
			endpoint:           "/v1/products/id/1",
			requestHeader:      "Bearer ",
			before: func() error {
				GetDetailProductHandler = func() (response dto.GetDetailProductResponse, err error) {
					return dto.GetDetailProductResponse{
						ID:          1,
						Sku:         "test",
						Name:        "test",
						Description: "test",
						Price:       1000,
						Stock:       10,
						Category:    "test",
						CategoryId:  1,
						ImageUrl:    "test.png",
						CreatedAt:   "2021-01-01T00:00:00Z",
						UpdatedAt:   "2021-01-01T00:00:00Z",
					}, nil
				}

				return nil
			},
		},
		{
			title:              "get detail product failed unauthorized",
			expectedErr:        middleware.ErrUnAuthorized,
			expectedValue:      dto.GetDetailProductResponse{},
			expectedStatusCode: fiber.StatusUnauthorized,
			endpoint:           "/v1/products/id/1",
			requestHeader:      "",
			before: func() error {
				GetDetailProductHandler = func() (response dto.GetDetailProductResponse, err error) {
					return dto.GetDetailProductResponse{}, middleware.ErrUnAuthorized
				}

				return middleware.ErrUnAuthorized
			},
		},
		{
			title:              "get detail product failed endpoint not found",
			expectedErr:        errors.New("endpoint not found"),
			expectedValue:      dto.GetDetailProductResponse{},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/product",
			requestHeader:      "Bearer ",
			before: func() error {
				GetDetailProductHandler = func() (response dto.GetDetailProductResponse, err error) {
					return dto.GetDetailProductResponse{}, errors.New("endpoint not found")
				}

				return errors.New("endpoint not found")
			},
		},
		{
			title:              "get detail product failed product not found",
			expectedErr:        entity.ErrProductNotFound,
			expectedValue:      dto.GetDetailProductResponse{},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/products/id/9999",
			requestHeader:      "Bearer ",
			before: func() error {
				GetDetailProductHandler = func() (response dto.GetDetailProductResponse, err error) {
					return dto.GetDetailProductResponse{}, entity.ErrProductNotFound
				}

				return entity.ErrProductNotFound
			},
		},
		{
			title:              "get detail product failed internal server error",
			expectedErr:        errors.New("internal server error"),
			expectedValue:      dto.GetDetailProductResponse{},
			expectedStatusCode: fiber.StatusInternalServerError,
			endpoint:           "/v1/products/id/9999'--",
			requestHeader:      "Bearer ",
			before: func() error {
				GetDetailProductHandler = func() (response dto.GetDetailProductResponse, err error) {
					return dto.GetDetailProductResponse{}, errors.New("internal server error")
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

			mockService := mockProductService{}
			handler := NewProductHandler(mockService)

			router.Get("/v1/products/id/:product_id", middleware.AuthMiddleware(), handler.GetDetailProduct)

			request := httptest.NewRequest(fiber.MethodGet, test.endpoint, nil)
			request.Header.Set("Authorization", test.requestHeader+signedToken)

			resp, _ := router.Test(request, 1)

			require.Equal(t, test.expectedStatusCode, resp.StatusCode)
			require.Equal(t, test.expectedErr, beforeErr)
		})
	}
}

func TestGetDetailProductUserPerspective(t *testing.T) {
	type testCase struct {
		title              string
		expectedErr        error
		expectedValue      dto.GetDetailProductUserPerspectiveResponse
		expectedStatusCode int
		endpoint           string
		requestHeader      string
		before             func() error
	}

	var testCases = []testCase{
		{
			title:       "get detail product user perspective success",
			expectedErr: nil,
			expectedValue: dto.GetDetailProductUserPerspectiveResponse{
				ID:          1,
				Sku:         "test",
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       10,
				Category:    "test",
				CategoryId:  1,
				Merchant: dto.Merchant{
					ID:   1,
					Name: "test",
					City: "test",
				},
				ImageUrl:  "test.png",
				CreatedAt: "2021-01-01T00:00:00Z",
				UpdatedAt: "2021-01-01T00:00:00Z",
			},
			expectedStatusCode: fiber.StatusOK,
			endpoint:           "/v1/products/detail/test",
			requestHeader:      "Bearer ",
			before: func() error {
				GetDetailProductUserPerspectiveHandler = func() (response dto.GetDetailProductUserPerspectiveResponse, err error) {
					return dto.GetDetailProductUserPerspectiveResponse{
						ID:          1,
						Sku:         "test",
						Name:        "test",
						Description: "test",
						Price:       1000,
						Stock:       10,
						Category:    "test",
						CategoryId:  1,
						Merchant: dto.Merchant{
							ID:   1,
							Name: "test",
							City: "test",
						},
						ImageUrl:  "test.png",
						CreatedAt: "2021-01-01T00:00:00Z",
						UpdatedAt: "2021-01-01T00:00:00Z",
					}, nil
				}

				return nil
			},
		},
		{
			title:              "get detail product user perspective failed unauthorized",
			expectedErr:        middleware.ErrUnAuthorized,
			expectedValue:      dto.GetDetailProductUserPerspectiveResponse{},
			expectedStatusCode: fiber.StatusUnauthorized,
			endpoint:           "/v1/products/detail/test",
			requestHeader:      "",
			before: func() error {
				GetDetailProductUserPerspectiveHandler = func() (response dto.GetDetailProductUserPerspectiveResponse, err error) {
					return dto.GetDetailProductUserPerspectiveResponse{}, middleware.ErrUnAuthorized
				}

				return middleware.ErrUnAuthorized
			},
		},
		{
			title:              "get detail product user perspective failed endpoint not found",
			expectedErr:        errors.New("endpoint not found"),
			expectedValue:      dto.GetDetailProductUserPerspectiveResponse{},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/products/details/test",
			requestHeader:      "Bearer ",
			before: func() error {
				GetDetailProductUserPerspectiveHandler = func() (response dto.GetDetailProductUserPerspectiveResponse, err error) {
					return dto.GetDetailProductUserPerspectiveResponse{}, errors.New("endpoint not found")
				}

				return errors.New("endpoint not found")
			},
		},
		{
			title:              "get detail product user perspective failed product by sku not found",
			expectedErr:        entity.ErrProductNotFound,
			expectedValue:      dto.GetDetailProductUserPerspectiveResponse{},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/products/detail/221213131kdnsaxx",
			requestHeader:      "Bearer ",
			before: func() error {
				GetDetailProductUserPerspectiveHandler = func() (response dto.GetDetailProductUserPerspectiveResponse, err error) {
					return dto.GetDetailProductUserPerspectiveResponse{}, entity.ErrProductNotFound
				}

				return entity.ErrProductNotFound
			},
		},
		{
			title:              "get detail product user perspective failed internal server error",
			expectedErr:        errors.New("internal server error"),
			expectedValue:      dto.GetDetailProductUserPerspectiveResponse{},
			expectedStatusCode: fiber.StatusInternalServerError,
			endpoint:           "/v1/products/detail/test'--",
			requestHeader:      "Bearer ",
			before: func() error {
				GetDetailProductUserPerspectiveHandler = func() (response dto.GetDetailProductUserPerspectiveResponse, err error) {
					return dto.GetDetailProductUserPerspectiveResponse{}, errors.New("internal server error")
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

			mockService := mockProductService{}
			handler := NewProductHandler(mockService)

			router.Get("/v1/products/detail/:sku", middleware.AuthMiddleware(), handler.GetDetailProductUserPerspective)

			request := httptest.NewRequest(fiber.MethodGet, test.endpoint, nil)
			request.Header.Set("Authorization", test.requestHeader+signedToken)

			resp, _ := router.Test(request, 1)

			require.Equal(t, test.expectedStatusCode, resp.StatusCode)
			require.Equal(t, test.expectedErr, beforeErr)
		})
	}
}

func TestGetListProductHandler(t *testing.T) {
	type testCase struct {
		title              string
		expectedErr        error
		expectedValue      []dto.GetListProductResponse
		expectedStatusCode int
		endpoint           string
		requestHeader      string
		before             func() error
	}

	var testCases = []testCase{
		{
			title:       "get list user success",
			expectedErr: nil,
			expectedValue: []dto.GetListProductResponse{
				{
					ID:          1,
					Sku:         "test",
					Name:        "test",
					Description: "test",
					Price:       1000,
					Stock:       10,
					Category:    "test",
					ImageUrl:    "test.png",
				},
				{
					ID:          2,
					Sku:         "test",
					Name:        "test",
					Description: "test",
					Price:       1000,
					Stock:       10,
					Category:    "test",
					ImageUrl:    "test.png",
				},
			},
			expectedStatusCode: fiber.StatusOK,
			endpoint:           "/v1/products",
			requestHeader:      "Bearer ",
			before: func() error {
				GetListProductHandler = func() (response []dto.GetListProductResponse, totalData int, err error) {
					return []dto.GetListProductResponse{
						{
							ID:          1,
							Sku:         "test",
							Name:        "test",
							Description: "test",
							Price:       1000,
							Stock:       10,
							Category:    "test",
							ImageUrl:    "test.png",
						},
						{
							ID:          2,
							Sku:         "test",
							Name:        "test",
							Description: "test",
							Price:       1000,
							Stock:       10,
							Category:    "test",
							ImageUrl:    "test.png",
						},
					}, 2, nil
				}

				return nil
			},
		},
		{
			title:              "get list product failed unauthorized",
			expectedErr:        middleware.ErrUnAuthorized,
			expectedValue:      []dto.GetListProductResponse{},
			expectedStatusCode: fiber.StatusUnauthorized,
			endpoint:           "/v1/products",
			requestHeader:      "",
			before: func() error {
				GetListProductHandler = func() (response []dto.GetListProductResponse, totalData int, err error) {
					return []dto.GetListProductResponse{}, 0, middleware.ErrUnAuthorized
				}

				return middleware.ErrUnAuthorized
			},
		},
		{
			title:              "get detail product user perspective failed endpoint not found",
			expectedErr:        errors.New("endpoint not found"),
			expectedValue:      []dto.GetListProductResponse{},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/product",
			requestHeader:      "Bearer ",
			before: func() error {
				GetListProductHandler = func() (response []dto.GetListProductResponse, totalData int, err error) {
					return []dto.GetListProductResponse{}, 0, errors.New("endpoint not found")
				}

				return errors.New("endpoint not found")
			},
		},
		{
			title:              "get list product failed internal server error",
			expectedErr:        errors.New("internal server error"),
			expectedValue:      []dto.GetListProductResponse{},
			expectedStatusCode: fiber.StatusInternalServerError,
			endpoint:           "/v1/products",
			requestHeader:      "Bearer ",
			before: func() error {
				GetListProductHandler = func() (response []dto.GetListProductResponse, totalData int, err error) {
					return []dto.GetListProductResponse{}, 0, errors.New("internal server error")
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

			mockService := mockProductService{}
			handler := NewProductHandler(mockService)

			router.Get("/v1/products", middleware.AuthMiddleware(), handler.GetListProduct)

			request := httptest.NewRequest(fiber.MethodGet, test.endpoint, nil)
			request.Header.Set("Authorization", test.requestHeader+signedToken)

			resp, _ := router.Test(request, 1)

			require.Equal(t, test.expectedStatusCode, resp.StatusCode)
			require.Equal(t, test.expectedErr, beforeErr)
		})
	}
}

func TestUpdateProductHandler(t *testing.T) {
	type testCase struct {
		title              string
		expectedErr        error
		request            dto.CreateOrUpdateProductRequest
		expectedStatusCode int
		endpoint           string
		contentType        string
		requestHeader      string
		before             func() error
	}

	var testCases = []testCase{
		{
			title:       "update product success",
			expectedErr: nil,
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusOK,
			endpoint:           "/v1/products/id/1",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return nil
				}

				return nil
			},
		},
		{
			title:       "update product failed unauthorized",
			expectedErr: middleware.ErrUnAuthorized,
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusUnauthorized,
			endpoint:           "/v1/products/id/:product_id",
			contentType:        "application/json",
			requestHeader:      "",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return middleware.ErrUnAuthorized
				}

				return middleware.ErrUnAuthorized
			},
		},
		{
			title:       "update product failed endpoint not found",
			expectedErr: errors.New("endpoint not found"),
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/product",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return errors.New("endpoint not found")
				}

				return errors.New("endpoint not found")
			},
		},
		{
			title:       "update product failed invalid content type",
			expectedErr: errors.New("invalid content type"),
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/product",
			contentType:        "application/xml",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return errors.New("invalid content type")
				}

				return errors.New("invalid content type")
			},
		},
		{
			title:       "update product failed name is required",
			expectedErr: entity.ErrProductNameIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				Name:        "",
				Description: "test",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products/id/1",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return entity.ErrProductNameIsRequired
				}

				return entity.ErrProductNameIsRequired
			},
		},
		{
			title:       "update product failed description is required",
			expectedErr: entity.ErrDescriptionIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "",
				Price:       1000,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products/id/1",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return entity.ErrDescriptionIsRequired
				}

				return entity.ErrDescriptionIsRequired
			},
		},
		{
			title:       "update product failed price is required",
			expectedErr: entity.ErrPriceIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "test",
				Price:       0,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products/id/1",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return entity.ErrPriceIsRequired
				}

				return entity.ErrPriceIsRequired
			},
		},
		{
			title:       "update product failed price is invalid",
			expectedErr: entity.ErrPriceIsInvalid,
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "test",
				Price:       -1,
				Stock:       100,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products/id/1",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return entity.ErrPriceIsInvalid
				}

				return entity.ErrPriceIsInvalid
			},
		},
		{
			title:       "update product failed stock is required",
			expectedErr: entity.ErrStockIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       0,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products/id/1",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return entity.ErrStockIsRequired
				}

				return entity.ErrStockIsRequired
			},
		},
		{
			title:       "update product failed stock is invalid",
			expectedErr: entity.ErrStockIsInvalid,
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       -1,
				CategoryId:  3,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products/id/1",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return entity.ErrStockIsInvalid
				}

				return entity.ErrStockIsInvalid
			},
		},
		{
			title:       "update product failed category id is required",
			expectedErr: entity.ErrCategoryIdIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       10,
				CategoryId:  0,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products/id/1",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return entity.ErrCategoryIdIsRequired
				}

				return entity.ErrCategoryIdIsRequired
			},
		},
		{
			title:       "update product failed image url is required",
			expectedErr: entity.ErrImageUrlIsRequired,
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       10,
				CategoryId:  1,
				ImageUrl:    "",
			},
			expectedStatusCode: fiber.StatusBadRequest,
			endpoint:           "/v1/products/id/1",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return entity.ErrImageUrlIsRequired
				}

				return entity.ErrImageUrlIsRequired
			},
		},
		{
			title:       "update product failed product not found",
			expectedErr: entity.ErrProductNotFound,
			request: dto.CreateOrUpdateProductRequest{
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       10,
				CategoryId:  1,
				ImageUrl:    "test.png",
			},
			expectedStatusCode: fiber.StatusNotFound,
			endpoint:           "/v1/products/id/999999",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
					return entity.ErrProductNotFound
				}

				return entity.ErrProductNotFound
			},
		},
		{
			title:       "update product failed internal server error",
			expectedErr: errors.New("internal server error"),
			request: dto.CreateOrUpdateProductRequest{
				ID:          -999999999999,
				Name:        "test",
				Description: "test",
				Price:       1000,
				Stock:       10,
				CategoryId:  1,
				ImageUrl:    "image.png",
			},
			expectedStatusCode: fiber.StatusInternalServerError,
			endpoint:           "/v1/products/id/1",
			contentType:        "application/json",
			requestHeader:      "Bearer ",
			before: func() error {
				UpdateProductHandler = func() (err error) {
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

			mockService := mockProductService{}
			handler := NewProductHandler(mockService)

			router.Post("/v1/products/id/:product_id", middleware.AuthMiddleware(), handler.UpdateProduct)

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
