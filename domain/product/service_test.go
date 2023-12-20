package product

import (
	"context"
	"errors"
	"testing"

	"github.com/ecommerce/dto"
	"github.com/ecommerce/entity"
	"github.com/stretchr/testify/require"
)

var svc = ProductService{}

type mockProductRepository struct{}
type mockMerchantRepository struct{}
type mockCategoryRepository struct{}

// Create implements Repository.
func (mockProductRepository) Create(ctx context.Context, product entity.Product) (err error) {
	return CreateProduct()
}

// GetById implements Repository.
func (mockProductRepository) GetById(ctx context.Context, id int) (product entity.Product, err error) {
	return GetProductById()
}

// GetByMerchantId implements Repository.
func (mockProductRepository) GetByMerchantId(ctx context.Context, queryParam string, limit int, page int, merchantId int) (products []entity.Product, totalData int, err error) {
	return GetProductByMerchantId()
}

// GetBySku implements Repository.
func (mockProductRepository) GetBySku(ctx context.Context, sku string) (product entity.Product, err error) {
	return GetProductBySku()
}

// Update implements Repository.
func (mockProductRepository) Update(ctx context.Context, product entity.Product) (err error) {
	return UpdateProduct()
}

// GetByCreatedBy implements merchant.Repository.
func (mockMerchantRepository) GetByCreatedBy(ctx context.Context, createdBy string) (merchant entity.Merchant, err error) {
	return GetMerchantByCreatedBy()
}

// Create implements category.Repository.
func (mockCategoryRepository) Create(ctx context.Context, category entity.Category) (err error) {
	return nil
}

// GetAll implements category.Repository.
func (mockCategoryRepository) GetAll(ctx context.Context) (categories []entity.Category, err error) {
	return nil, nil
}

// GetById implements category.Repository.
func (mockCategoryRepository) GetById(ctx context.Context, id int) (category entity.Category, err error) {
	return GetCategoryById()
}

var (
	CreateProduct          func() (err error)
	GetProductById         func() (product entity.Product, err error)
	GetProductByMerchantId func() (products []entity.Product, totalData int, err error)
	GetProductBySku        func() (product entity.Product, err error)
	UpdateProduct          func() (err error)
	GetMerchantByCreatedBy func() (merchant entity.Merchant, err error)
	GetCategoryById        func() (category entity.Category, err error)
)

func init() {
	mockProduct := mockProductRepository{}
	mockMerchant := mockMerchantRepository{}
	mockCategory := mockCategoryRepository{}

	svc = NewProductService(mockProduct, mockMerchant, mockCategory)
}

func TestCreateProduct(t *testing.T) {
	type testCase struct {
		title       string
		expectedErr error
		request     entity.Product
		before      func()
	}

	var testCases = []testCase{
		{
			title:       "create product success",
			expectedErr: nil,
			request: entity.Product{
				ID:          1,
				Name:        "product 1",
				Description: "description",
				Price:       10000,
				Stock:       10,
				CategoryId:  1,
				ImageUrl:    "image.png",
			},
			before: func() {
				GetMerchantByCreatedBy = func() (merchant entity.Merchant, err error) {
					return entity.Merchant{
						ID:        1,
						Name:      "merchant 1",
						Role:      "merchant",
						CreatedBy: "1",
					}, nil
				}

				GetCategoryById = func() (category entity.Category, err error) {
					return entity.Category{
						ID:   1,
						Name: "category 1",
					}, nil
				}

				CreateProduct = func() (err error) {
					return nil
				}
			},
		},
		{
			title:       "create product failed merchant not found",
			expectedErr: errors.New("merchant not found"),
			request: entity.Product{
				ID:          1,
				Name:        "product 1",
				Description: "description",
				Price:       10000,
				Stock:       10,
				CategoryId:  1,
				ImageUrl:    "image.png",
			},
			before: func() {
				GetMerchantByCreatedBy = func() (merchant entity.Merchant, err error) {
					return entity.Merchant{}, errors.New("merchant not found")
				}

				GetCategoryById = func() (category entity.Category, err error) {
					return entity.Category{
						ID:   1,
						Name: "category 1",
					}, nil
				}

				CreateProduct = func() (err error) {
					return nil
				}
			},
		},
		{
			title:       "create product failed invalid role",
			expectedErr: entity.ErrInvalidRole,
			request: entity.Product{
				ID:          1,
				Name:        "product 1",
				Description: "description",
				Price:       10000,
				Stock:       10,
				CategoryId:  1,
				ImageUrl:    "image.png",
			},
			before: func() {
				GetMerchantByCreatedBy = func() (merchant entity.Merchant, err error) {
					return entity.Merchant{
						ID:        1,
						Name:      "merchant 1",
						Role:      "user",
						CreatedBy: "1",
					}, nil
				}

				GetCategoryById = func() (category entity.Category, err error) {
					return entity.Category{
						ID:   1,
						Name: "category 1",
					}, nil
				}

				CreateProduct = func() (err error) {
					return nil
				}
			},
		},
		{
			title:       "create product failed category not found",
			expectedErr: entity.ErrCategoryNotFound,
			request: entity.Product{
				ID:          1,
				Name:        "product 1",
				Description: "description",
				Price:       10000,
				Stock:       10,
				CategoryId:  99999,
				ImageUrl:    "image.png",
			},
			before: func() {
				GetMerchantByCreatedBy = func() (merchant entity.Merchant, err error) {
					return entity.Merchant{
						ID:        1,
						Name:      "merchant 1",
						Role:      "merchant",
						CreatedBy: "1",
					}, nil
				}

				GetCategoryById = func() (category entity.Category, err error) {
					return entity.Category{}, entity.ErrCategoryNotFound
				}

				CreateProduct = func() (err error) {
					return nil
				}
			},
		},
		{
			title:       "create product failed internal server error",
			expectedErr: errors.New("internal server error"),
			request: entity.Product{
				ID:          1,
				Name:        "product 1",
				Description: "description",
				Price:       10000,
				Stock:       10,
				CategoryId:  1,
				ImageUrl:    "image.png",
			},
			before: func() {
				GetMerchantByCreatedBy = func() (merchant entity.Merchant, err error) {
					return entity.Merchant{
						ID:        1,
						Name:      "merchant 1",
						Role:      "merchant",
						CreatedBy: "1",
					}, nil
				}

				GetCategoryById = func() (category entity.Category, err error) {
					return entity.Category{
						ID:   1,
						Name: "category 1",
					}, nil
				}

				CreateProduct = func() (err error) {
					return errors.New("internal server error")
				}
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.before()

			err := svc.CreateProduct(context.Background(), test.request, "1")
			require.Equal(t, test.expectedErr, err)
		})
	}
}

func TestGetDetailoProduct(t *testing.T) {
	type testCase struct {
		title         string
		expectedErr   error
		expectedValue dto.GetDetailProductResponse
		before        func()
	}

	var testCases = []testCase{
		{
			title:       "get product by id success",
			expectedErr: nil,
			expectedValue: dto.GetDetailProductResponse{
				ID:          1,
				Name:        "product 1",
				Description: "description",
				Price:       10000,
				Stock:       10,
				CategoryId:  1,
				ImageUrl:    "image.png",
				Sku:         "sku",
				Category:    "category 1",
			},
			before: func() {
				GetMerchantByCreatedBy = func() (merchant entity.Merchant, err error) {
					return entity.Merchant{
						ID:        1,
						Name:      "merchant 1",
						Role:      "merchant",
						CreatedBy: "1",
					}, nil
				}

				GetProductById = func() (product entity.Product, err error) {
					return entity.Product{
						ID:           1,
						Name:         "product 1",
						Description:  "description",
						Price:        10000,
						Stock:        10,
						CategoryId:   1,
						MerchantId:   1,
						ImageUrl:     "image.png",
						Sku:          "sku",
						Category:     "category 1",
						MerchantName: "merchant 1",
						MerchantCity: "city 1",
					}, nil
				}
			},
		},
		{
			title:         "get product by id failed merchant not found",
			expectedErr:   errors.New("merchant not found"),
			expectedValue: dto.GetDetailProductResponse{},
			before: func() {
				GetMerchantByCreatedBy = func() (merchant entity.Merchant, err error) {
					return entity.Merchant{}, errors.New("merchant not found")
				}

				GetProductById = func() (product entity.Product, err error) {
					return entity.Product{}, errors.New("merchant not found")
				}
			},
		},
		{
			title:         "get product by id failed invalid role",
			expectedErr:   entity.ErrInvalidRole,
			expectedValue: dto.GetDetailProductResponse{},
			before: func() {
				GetMerchantByCreatedBy = func() (merchant entity.Merchant, err error) {
					return entity.Merchant{
						ID:        1,
						Name:      "merchant 1",
						Role:      "user",
						CreatedBy: "1",
					}, nil
				}

				GetProductById = func() (product entity.Product, err error) {
					return entity.Product{}, errors.New("merchant not found")
				}
			},
		},
		{
			title:         "get product by id failed product not found",
			expectedErr:   entity.ErrProductNotFound,
			expectedValue: dto.GetDetailProductResponse{},
			before: func() {
				GetMerchantByCreatedBy = func() (merchant entity.Merchant, err error) {
					return entity.Merchant{
						ID:        1,
						Name:      "merchant 1",
						Role:      "merchant",
						CreatedBy: "1",
					}, nil
				}

				GetProductById = func() (product entity.Product, err error) {
					return entity.Product{}, entity.ErrProductNotFound
				}
			},
		},
		{
			title:         "get product by id failed internal server error",
			expectedErr:   errors.New("internal server error"),
			expectedValue: dto.GetDetailProductResponse{},
			before: func() {
				GetMerchantByCreatedBy = func() (merchant entity.Merchant, err error) {
					return entity.Merchant{
						ID:        1,
						Name:      "merchant 1",
						Role:      "merchant",
						CreatedBy: "1",
					}, nil
				}

				GetProductById = func() (product entity.Product, err error) {
					return entity.Product{}, errors.New("internal server error")
				}
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.before()

			product, err := svc.GetDetailProduct(context.Background(), 1, "1")
			require.Equal(t, test.expectedErr, err)
			require.Equal(t, test.expectedValue, product)
		})
	}
}
