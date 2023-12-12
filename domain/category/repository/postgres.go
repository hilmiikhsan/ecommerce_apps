package repository

import (
	"context"

	"github.com/ecommerce/entity"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return CategoryRepository{
		db: db,
	}
}

func (c CategoryRepository) Create(ctx context.Context, category entity.Category) (err error) {
	stmt, err := c.db.PrepareNamedContext(ctx, queryCreate)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, category)
	if err != nil {
		return
	}

	return
}

func (c CategoryRepository) GetAll(ctx context.Context) (categories []entity.Category, err error) {
	err = c.db.SelectContext(ctx, &categories, queryGetAll)
	if err != nil {
		return
	}

	return
}
