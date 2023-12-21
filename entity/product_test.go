package entity

import (
	"testing"

	"github.com/ecommerce/dto"
	"github.com/stretchr/testify/require"
)

func TestEntityProduct(t *testing.T) {
	t.Run("err : product name is required", func(t *testing.T) {
		var req = dto.CreateOrUpdateProductRequest{
			Name:        "",
			Description: "test",
			Price:       1000,
			Stock:       10,
			CategoryId:  1,
			ImageUrl:    "test",
		}

		_, err := NewProduct().Validate(req, "1")
		require.NotNil(t, err)
		require.Equal(t, ErrProductNameIsRequired, err)
	})

	t.Run("err : product description is required", func(t *testing.T) {
		var req = dto.CreateOrUpdateProductRequest{
			Name:        "test",
			Description: "",
			Price:       1000,
			Stock:       10,
			CategoryId:  1,
			ImageUrl:    "test",
		}

		_, err := NewProduct().Validate(req, "1")
		require.NotNil(t, err)
		require.Equal(t, ErrDescriptionIsRequired, err)
	})

	t.Run("err : product price is required", func(t *testing.T) {
		var req = dto.CreateOrUpdateProductRequest{
			Name:        "test",
			Description: "test",
			Price:       0,
			Stock:       10,
			CategoryId:  1,
			ImageUrl:    "test",
		}

		_, err := NewProduct().Validate(req, "1")
		require.NotNil(t, err)
		require.Equal(t, ErrPriceIsRequired, err)
	})

	t.Run("err : product price is invalid", func(t *testing.T) {
		var req = dto.CreateOrUpdateProductRequest{
			Name:        "test",
			Description: "test",
			Price:       -1,
			Stock:       10,
			CategoryId:  1,
			ImageUrl:    "test",
		}

		_, err := NewProduct().Validate(req, "1")
		require.NotNil(t, err)
		require.Equal(t, ErrPriceIsInvalid, err)
	})

	t.Run("err : product stock is required", func(t *testing.T) {
		var req = dto.CreateOrUpdateProductRequest{
			Name:        "test",
			Description: "test",
			Price:       1000,
			Stock:       0,
			CategoryId:  1,
			ImageUrl:    "test",
		}

		_, err := NewProduct().Validate(req, "1")
		require.NotNil(t, err)
		require.Equal(t, ErrStockIsRequired, err)
	})

	t.Run("err : product stock is invalid", func(t *testing.T) {
		var req = dto.CreateOrUpdateProductRequest{
			Name:        "test",
			Description: "test",
			Price:       1000,
			Stock:       -1,
			CategoryId:  1,
			ImageUrl:    "test",
		}

		_, err := NewProduct().Validate(req, "1")
		require.NotNil(t, err)
		require.Equal(t, ErrStockIsInvalid, err)
	})

	t.Run("err : product category id is required", func(t *testing.T) {
		var req = dto.CreateOrUpdateProductRequest{
			Name:        "test",
			Description: "test",
			Price:       1000,
			Stock:       10,
			CategoryId:  0,
			ImageUrl:    "test",
		}

		_, err := NewProduct().Validate(req, "1")
		require.NotNil(t, err)
		require.Equal(t, ErrCategoryIdIsRequired, err)
	})

	t.Run("err : product image url is required", func(t *testing.T) {
		var req = dto.CreateOrUpdateProductRequest{
			Name:        "test",
			Description: "test",
			Price:       1000,
			Stock:       10,
			CategoryId:  1,
			ImageUrl:    "",
		}

		_, err := NewProduct().Validate(req, "1")
		require.NotNil(t, err)
		require.Equal(t, ErrImageUrlIsRequired, err)
	})

	t.Run("success : validate product request", func(t *testing.T) {
		var req = dto.CreateOrUpdateProductRequest{
			Name:        "test",
			Description: "test",
			Price:       1000,
			Stock:       10,
			CategoryId:  1,
			ImageUrl:    "test",
		}

		_, err := NewProduct().Validate(req, "1")
		require.Nil(t, err)
	})

	t.Run("err : invalid role", func(t *testing.T) {
		err := NewProduct().CheckUserRole("user")
		require.NotNil(t, err)
		require.Equal(t, ErrInvalidRole, err)
	})

	t.Run("success : valid role", func(t *testing.T) {
		err := NewProduct().CheckUserRole("merchant")
		require.Nil(t, err)
	})
}
