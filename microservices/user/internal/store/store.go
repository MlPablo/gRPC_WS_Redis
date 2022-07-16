package store

import (
	"os"

	"github.com/go-redis/redis/v9"
)

type Storage interface {
	NewCRUD() CRUD
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

func (s *storage) NewCRUD() CRUD {
	return &crud{s}
}
