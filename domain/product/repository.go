package product

import (
	"context"

	"github.com/ecommerce/entity"
)

type Repository interface {
	Create(ctx context.Context, product entity.Product) (err error)
	GetByMerchantId(ctx context.Context, queryParam string, limit, page, merchantId int) (products []entity.Product, totalData int, err error)
	GetById(ctx context.Context, id int) (product entity.Product, err error)
	Update(ctx context.Context, product entity.Product) (err error)
	GetBySku(ctx context.Context, sku string) (product entity.Product, err error)
}
