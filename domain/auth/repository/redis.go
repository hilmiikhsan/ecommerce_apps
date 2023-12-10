package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type RedisRepository struct {
	redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) RedisRepository {
	return RedisRepository{
		redis: redis,
	}
}

func (r RedisRepository) Set(ctx context.Context, timeLimit int, token, id, email string) (err error) {
	key := fmt.Sprintf("%s:%s", email, id)
	ttl := time.Duration(timeLimit) * time.Hour
	err = r.redis.Set(ctx, key, token, ttl).Err()
	if err != nil {
		logrus.Error(err)
		return
	}

	return
}

func (r RedisRepository) Get(ctx context.Context, id, email string) (token string, err error) {
	key := fmt.Sprintf("%s:%s", email, id)
	token, err = r.redis.Get(ctx, key).Result()
	if err != nil {
		return "", nil
	}

	return token, nil
}
