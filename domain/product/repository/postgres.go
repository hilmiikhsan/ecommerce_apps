package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ecommerce/entity"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return ProductRepository{
		db: db,
	}
}

func (p ProductRepository) Create(ctx context.Context, product entity.Product) (err error) {
	stmt, err := p.db.PrepareNamedContext(ctx, queryCreate)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, product)
	if err != nil {
		return
	}

	return
}

func (p ProductRepository) GetByMerchantId(ctx context.Context, queryParam string, limit, page, merchantId int) (products []entity.Product, totalData int, err error) {
	offset := (page - 1) * limit
	filter := mappingQueryFilter(queryParam)
	queryLimitOffset := fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
	query := fmt.Sprintf("%s %s %s", queryGetByMerchantId, filter, queryLimitOffset)
	queryCount := fmt.Sprintf("%s %s %s", queryCountByMerchantId, filter, queryLimitOffset)

	err = p.db.SelectContext(ctx, &products, query, merchantId)
	if err != nil {
		return
	}

	err = p.db.GetContext(ctx, &totalData, queryCount, merchantId)
	if err != nil {
		if err == sql.ErrNoRows {
			totalData = 0
			return []entity.Product{}, totalData, nil
		}
		return
	}

	return
}

func (p ProductRepository) GetById(ctx context.Context, id int) (product entity.Product, err error) {
	err = p.db.GetContext(ctx, &product, queryGetById, id)
	if err != nil {
		return
	}

	return
}

func (p ProductRepository) Update(ctx context.Context, product entity.Product) (err error) {
	stmt, err := p.db.PrepareNamedContext(ctx, queryUpdate)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, product)
	if err != nil {
		return
	}

	return
}

func (p ProductRepository) GetBySku(ctx context.Context, sku string) (product entity.Product, err error) {
	err = p.db.GetContext(ctx, &product, queryGetBySku, sku)
	if err != nil {
		return
	}

	return
}

func mappingQueryFilter(queryParam string) string {
	filter := ""

	if queryParam != "" {
		filter = fmt.Sprintf("%s AND p.name ILIKE '%%%s%%'", filter, queryParam)
	}

	return filter
}
