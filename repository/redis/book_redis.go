package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type bookRepositoryRedis struct {
	rdb *redis.Client
}

var Ctx = context.Background()

func NewBookRepositoryRedis(rdb *redis.Client) *bookRepositoryRedis {
	return &bookRepositoryRedis{rdb: rdb}
}

func (r *bookRepositoryRedis) Set(data BookRedis) error {
	err := r.rdb.Set(Ctx, data.Key, data.Value, data.Expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *bookRepositoryRedis) Get(key string) (string, error) {
	val, err := r.rdb.Get(Ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *bookRepositoryRedis) Delete(key string) error {
	err := r.rdb.Del(Ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
