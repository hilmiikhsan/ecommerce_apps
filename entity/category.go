package entity

import (
	"errors"

	"github.com/ecommerce/dto"
)

var (
	ErrCategoryNameIsRequired = errors.New("category name is required")
)

type Category struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	CreatedBy string `db:"created_by"`
}

func NewCategory() Category {
	return Category{}
}

func (ca Category) Validate(req dto.CreateCategoryRequest, id string) (Category, error) {
	if req.Name == "" {
		return ca, ErrCategoryNameIsRequired
	}

	ca.Name = req.Name
	ca.CreatedBy = id

	return ca, nil
}

func (ca Category) CategoryResponse(categories []Category) []dto.GetListCategoryResponse {
	responses := []dto.GetListCategoryResponse{}

	for _, category := range categories {
		response := dto.GetListCategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		}

		responses = append(responses, response)
	}

	return responses
}
