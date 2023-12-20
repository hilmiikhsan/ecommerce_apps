package category

import (
	"context"
	"errors"
	"testing"

	"github.com/ecommerce/dto"
	"github.com/ecommerce/entity"
	"github.com/stretchr/testify/require"
)

var svc = CategoryService{}

type mockCategoryRepository struct{}

// GetById implements Repository.
func (mockCategoryRepository) GetById(ctx context.Context, id int) (category entity.Category, err error) {
	return category, nil
}

// Create implements Repository.
func (mockCategoryRepository) Create(ctx context.Context, category entity.Category) (err error) {
	return Create()
}

// GetAll implements Repository.
func (mockCategoryRepository) GetAll(ctx context.Context) (categories []entity.Category, err error) {
	return GetAll()
}

var (
	Create func() (err error)
	GetAll func() (categories []entity.Category, err error)
)

func init() {
	mock := mockCategoryRepository{}

	svc = NewCategoryService(mock)
}

func TestCreateCategory(t *testing.T) {
	type testCase struct {
		title       string
		expectedErr error
		request     entity.Category
		before      func()
	}

	var testCases = []testCase{
		{
			title:       "create category success",
			expectedErr: nil,
			request: entity.Category{
				ID:   1,
				Name: "category 1",
			},
			before: func() {
				Create = func() (err error) {
					return nil
				}
			},
		},
		{
			title:       "create category failed internal server error",
			expectedErr: errors.New("internal server error"),
			request: entity.Category{
				ID:   1,
				Name: "category 1",
			},
			before: func() {
				Create = func() (err error) {
					return errors.New("internal server error")
				}
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.before()

			err := svc.CreateCategory(context.Background(), test.request)
			require.Equal(t, test.expectedErr, err)
		})
	}
}

func TestGetListCategory(t *testing.T) {
	type testCase struct {
		title         string
		expectedErr   error
		expectedValue []dto.GetListCategoryResponse
		before        func()
	}

	var testCases = []testCase{
		{
			title:       "get all category success",
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
			before: func() {
				GetAll = func() (categories []entity.Category, err error) {
					return []entity.Category{
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
			},
		},
		{
			title:         "get all category failed internal server error",
			expectedErr:   errors.New("internal server error"),
			expectedValue: []dto.GetListCategoryResponse{},
			before: func() {
				GetAll = func() (categories []entity.Category, err error) {
					return nil, errors.New("internal server error")
				}
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.before()

			_, err := svc.GetListCategory(context.Background())
			require.Equal(t, test.expectedErr, err)
		})
	}
}
