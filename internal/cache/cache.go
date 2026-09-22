package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/fastscripts/grpctest/internal/config"
	"github.com/redis/go-redis/v9"
)

type CacheService interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
	Ping(ctx context.Context) error
	Close() error
}

type RedisCache struct {
	client *redis.Client
}

type Opts struct {
	Config *config.Redis
}

func NewRedisCache(ctx context.Context, opts *Opts) (CacheService, error) {
	cfg := opts.Config
	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  time.Duration(cfg.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WithTimeout) * time.Second,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &RedisCache{client: rdb}, nil

}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisCache) Get(ctx context.Context, key string) (interface{}, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // Key does not exist
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *RedisCache) Close() error {
	if r == nil || r.client == nil {
		return fmt.Errorf("cannot close: redis is not initialized")
	}
	return r.client.Close()
}
