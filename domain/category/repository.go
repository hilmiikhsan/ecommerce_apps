package category

import (
	"context"

	"github.com/ecommerce/entity"
)

type Repository interface {
	Create(ctx context.Context, category entity.Category) (err error)
	GetAll(ctx context.Context) (categories []entity.Category, err error)
	GetById(ctx context.Context, id int) (category entity.Category, err error)
}
