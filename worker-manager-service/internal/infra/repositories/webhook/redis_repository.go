package webhook

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisRepository struct {
	client *redis.Client
	ctx    context.Context
	logger *zap.Logger
}

func NewRedisRepository(dsn string, ctx context.Context, logger *zap.Logger) (*RedisRepository, error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(opt)
	return &RedisRepository{
		client: rdb,
		ctx:    ctx,
		logger: logger,
	}, nil
}

func (r *RedisRepository) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (r *RedisRepository) Set(key string, value string) error {
	return r.client.Set(r.ctx, key, value, time.Minute*30).Err()
}

func (r *RedisRepository) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
