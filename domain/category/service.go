package category

import (
	"context"

	"github.com/ecommerce/dto"
	"github.com/ecommerce/entity"
)

type Service interface {
	CreateCategory(ctx context.Context, req entity.Category) (err error)
	GetListCategory(ctx context.Context) (response []dto.GetListCategoryResponse, err error)
}

type CategoryService struct {
	repository Repository
}

func NewCategoryService(repository Repository) CategoryService {
	return CategoryService{
		repository: repository,
	}
}

func (c CategoryService) CreateCategory(ctx context.Context, req entity.Category) (err error) {
	if err = c.repository.Create(ctx, req); err != nil {
		return
	}

	return
}

func (c CategoryService) GetListCategory(ctx context.Context) (response []dto.GetListCategoryResponse, err error) {
	categories, err := c.repository.GetAll(ctx)
	if err != nil {
		return
	}

	response = entity.NewCategory().CategoryResponse(categories)

	return
}
