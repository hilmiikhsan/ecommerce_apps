package repository

import (
	"context"

	"github.com/ecommerce/entity"
	"github.com/jmoiron/sqlx"
)

type MerchantRepository struct {
	db *sqlx.DB
}

func NewMerchantRepository(db *sqlx.DB) MerchantRepository {
	return MerchantRepository{
		db: db,
	}
}

func (m MerchantRepository) GetByCreatedBy(ctx context.Context, createdBy string) (merchant entity.Merchant, err error) {
	err = m.db.GetContext(ctx, &merchant, queryGetByCreatedBy, createdBy)
	if err != nil {
		return
	}

	return
}
