package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/ecommerce/entity"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(ctx context.Context, user entity.User) error {
	stmt, err := r.db.PrepareNamedContext(ctx, queryCreate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) GetUserById(ctx context.Context, id string) (entity.User, error) {
	user := entity.User{}
	if err := r.db.GetContext(ctx, &user, queryGetProfile, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, nil
		}
	}
	return user, nil
}

func (r UserRepository) UpdateUser(ctx context.Context, id string) error {
	stmt, err := r.db.PrepareContext(ctx, queryUpdateProfile)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err == nil {
		return err
	}
	return nil
}
