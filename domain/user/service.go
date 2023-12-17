package user

import (
	"context"

	"github.com/ecommerce/domain/auth"
	"github.com/ecommerce/entity"
)

type UserService struct {
	repository     Repository
	authRepository auth.Repository
}

func NewUserService(repository Repository, authRepository auth.Repository) UserService {
	return UserService{
		repository:     repository,
		authRepository: authRepository,
	}
}

func (u UserService) CreateUserProfile(ctx context.Context, req entity.User, email string) error {
	userAuth, err := u.authRepository.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if err = entity.NewUser().CheckUserRole(userAuth.Role); err != nil {
		return nil
	}

	err = u.repository.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) GetUser(ctx context.Context, req entity.User, token string) error {
	return nil
}
