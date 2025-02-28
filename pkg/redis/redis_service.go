package redispkg

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	REDIS_TTL = 15 * time.Minute
)

type RedisService interface {
	Set(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) (string, error)
}

type redisServiceImpl struct {
	client *redis.Client
}

func NewRedisService() RedisService {
	client := Client
	return &redisServiceImpl{client}
}

// Set implements RedisService.
func (svc *redisServiceImpl) Set(ctx context.Context, key string, value []byte) error {
	result, err := svc.client.Set(ctx, key, value, 0).Result()
	log.Printf("Redis set result: %v\n", result)
	return err
}

func (svc *redisServiceImpl) Get(ctx context.Context, key string) (string, error) {
	result, err := svc.client.Get(ctx, key).Result()
	log.Printf("Get result: %v\n", result)
	return result, err
}
