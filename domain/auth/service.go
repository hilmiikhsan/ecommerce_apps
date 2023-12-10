package auth

import (
	"context"

	"github.com/ecommerce/config"
	"github.com/ecommerce/entity"
)

type AuthService struct {
	repository Repository
	redis      RedisRepository
	cfg        config.JWT
}

func NewAuthService(repository Repository, redis RedisRepository, cfg config.JWT) AuthService {
	return AuthService{
		repository: repository,
		redis:      redis,
		cfg:        cfg,
	}
}

func (a AuthService) Register(ctx context.Context, req entity.Auth) (err error) {
	user, err := a.repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return
	}

	if err = req.CheckRequestEmail(req.Email, user.Email); err != nil {
		return
	}

	if err = req.EncryptPassword(); err != nil {
		return
	}

	if err = a.repository.Create(ctx, req); err != nil {
		return
	}

	return
}

func (a AuthService) Login(ctx context.Context, req entity.Auth) (response entity.Auth, accessToken string, err error) {
	user, err := a.repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return
	}

	if err = req.CheckRegisteredEmail(req.Email, user.Email); err != nil {
		return
	}

	if !req.ValidatePasswordFromPlainText(req.Password, user.Password) {
		return response, accessToken, entity.ErrInvalidEmailOrPassword
	}

	accessToken, err = a.redis.Get(ctx, user.ID, user.Email)
	if err != nil {
		return
	}

	if accessToken == "" {
		accessToken, err = user.GenerateAccessToken(accessToken)
		if err != nil {
			return
		}

		if err = a.redis.Set(ctx, a.cfg.TokenLifeTimeHour, accessToken, user.ID, user.Email); err != nil {
			return
		}
	}

	return user, accessToken, nil
}

func (a AuthService) UpdateRole(ctx context.Context, id, email string) (err error) {
	user, err := a.repository.GetByEmail(ctx, email)
	if err != nil {
		return
	}

	if err = entity.NewAuth().ValidateUserRole(user.Role); err != nil {
		return
	}

	if err = a.repository.UpdateRole(ctx, id); err != nil {
		return
	}

	return
}
