package redis

import (
	"context"
	"fmt"

	goredis "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*goredis.Client
}

func New(ctx context.Context, cfg RedisConfig) (*RedisClient, error) {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		_ = rdb.Close()
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}

	return &RedisClient{Client: rdb}, nil
}
