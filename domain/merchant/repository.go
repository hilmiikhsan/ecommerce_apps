package merchant

import (
	"context"

	"github.com/ecommerce/entity"
)

type Repository interface {
	GetByCreatedBy(ctx context.Context, createdBy string) (merchant entity.Merchant, err error)
}
