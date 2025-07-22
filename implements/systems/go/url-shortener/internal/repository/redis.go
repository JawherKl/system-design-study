package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

type RedisRepo struct {
	client *redis.Client
}

var ctx = context.Background()

func NewRedisRepo() *RedisRepo {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
	return &RedisRepo{client: client}
}

func (r *RedisRepo) Set(hash string, original string) error {
	return r.client.Set(ctx, hash, original, 0).Err()
}

func (r *RedisRepo) Get(hash string) (string, error) {
	return r.client.Get(ctx, hash).Result()
}