package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ecommerce/entity"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return AuthRepository{
		db: db,
	}
}

const (
	RoleMerchant = "merchant"
)

func (r AuthRepository) Create(ctx context.Context, user entity.Auth) (err error) {
	stmt, err := r.db.PrepareNamedContext(ctx, queryCreate)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user)
	if err != nil {
		return
	}

	return
}

func (r AuthRepository) GetByEmail(ctx context.Context, email string) (user entity.Auth, err error) {
	err = r.db.GetContext(ctx, &user, queryGetByEmail, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Auth{}, nil
		}
	}

	return
}

func (r AuthRepository) UpdateRole(ctx context.Context, id string) (err error) {
	stmt, err := r.db.PreparexContext(ctx, queryUpdateRole)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, RoleMerchant, id)
	if err != nil {
		return
	}

	return
}
