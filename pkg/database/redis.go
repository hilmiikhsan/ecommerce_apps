package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ecommerce/config"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func ConnectRedis(ctx context.Context, cfg config.Redis) (*redis.Client, error) {
	timeout := time.Duration(cfg.Timeout) * time.Second

	rdb := redis.NewClient(&redis.Options{
		Addr:        cfg.Addr,
		Password:    cfg.Password,
		DialTimeout: timeout,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Error("cannot connect to redis")
		return nil, err
	}

	fmt.Printf("success connect to redis %s", rdb)
	return rdb, nil
}
