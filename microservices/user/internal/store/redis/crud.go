package redisDB

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis/v9"

	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/models"
)

type redisDB struct {
	store *redis.Client
}

func New() *redisDB {
	return &redisDB{redis.NewClient(
		&redis.Options{
			Addr:     os.Getenv("REDIS_HOST") + os.Getenv("REDIS_URL"),
			Password: "",
			DB:       0,
		})}
}

func (c *redisDB) Create(ctx context.Context, user models.User) error {
	cont, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	if exist, _ := c.store.Get(cont, user.User).Result(); exist != "" {
		return errors.New("already exists")
	}

	if err := c.store.Set(context.Background(), user.User, user.Password, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (c *redisDB) Read(ctx context.Context, id string) (string, error) {
	cont, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	user := c.store.Get(cont, id)
	if user.Err() != nil {
		return "", user.Err()
	}
	return user.Val(), nil
}

func (c *redisDB) Update(ctx context.Context, user models.User) error {
	cont, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	if err := c.store.Get(cont, user.User).Err(); err != nil {
		return errors.New("no such user")
	}
	if err := c.store.GetSet(context.Background(), user.User, user.Password).Err(); err != nil {
		return err
	}
	return nil
}

func (c *redisDB) Delete(ctx context.Context, id string) error {
	cont, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	if err := c.store.GetDel(cont, id).Err(); err != nil {
		return err
	}
	return nil
}
