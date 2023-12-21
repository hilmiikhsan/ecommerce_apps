package entity

import (
	"testing"

	"github.com/ecommerce/dto"
	"github.com/stretchr/testify/require"
)

func TestEntityCategory(t *testing.T) {
	t.Run("err : category name is required", func(t *testing.T) {
		var req = dto.CreateCategoryRequest{
			Name: "",
		}

		_, err := NewCategory().Validate(req, "1")
		require.NotNil(t, err)
		require.Equal(t, ErrCategoryNameIsRequired, err)
	})

	t.Run("success : validate category request", func(t *testing.T) {
		var req = dto.CreateCategoryRequest{
			Name: "test",
		}

		_, err := NewCategory().Validate(req, "1")
		require.Nil(t, err)
	})
}
