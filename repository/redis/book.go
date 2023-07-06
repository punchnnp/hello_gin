package redis

import "time"

type BookRedis struct {
	Key        int
	Value      interface{}
	Expiration time.Duration
}

type BookRepositoryRedis interface {
	Set(BookRedis) error
	Get(int) (string, error)
	Delete(int) error
}
