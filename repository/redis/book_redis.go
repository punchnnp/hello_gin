package redis

import (
	"context"
	"fmt"

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
	err := r.rdb.Set(Ctx, fmt.Sprintf("%v", data.Key), data.Value, data.Expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *bookRepositoryRedis) Get(key int) (string, error) {
	val, err := r.rdb.Get(Ctx, fmt.Sprintf("%v", key)).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *bookRepositoryRedis) Delete(key int) error {
	err := r.rdb.Del(Ctx, fmt.Sprintf("%v", key)).Err()
	if err != nil {
		return err
	}
	return nil
}
