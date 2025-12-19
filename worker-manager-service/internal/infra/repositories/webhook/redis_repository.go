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

func (r *RedisRepository) Scan() ([]string, error) {
	var cursor uint64
	var keys []string

	for {
		batch, nextCursor, err := r.client.Scan(r.ctx, cursor, "*", 100).Result()
		if err != nil {
			return nil, err
		}

		keys = append(keys, batch...)
		cursor = nextCursor

		if cursor == 0 {
			break
		}
	}
	return keys, nil
}

func (r *RedisRepository) Set(key string, value string) error {
	return r.client.Set(r.ctx, key, value, time.Minute*30).Err()
}

func (r *RedisRepository) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
