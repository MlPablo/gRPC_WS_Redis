package store

import (
	"os"

	"github.com/go-redis/redis/v9"
)

type Storage interface {
	NewOrders() Orders
}

func New() Storage {
	return &storage{redis.NewClient(
		&redis.Options{
			Addr:     os.Getenv("REDIS_HOST") + os.Getenv("REDIS_URL"),
			Password: "",
			DB:       0,
		})}
}

type storage struct {
	store *redis.Client
}

func (s *storage) NewOrders() Orders {
	return &orders{s}
}
