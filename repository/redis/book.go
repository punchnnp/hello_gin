package redis

import "time"

type BookRedis struct {
	Key        string
	Value      string
	Expiration time.Duration
}

type BookRepositoryRedis interface {
	Set(BookRedis) error
	Get(string) (string, error)
	Delete(string) error
}
