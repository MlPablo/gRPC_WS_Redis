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
	Store *redis.Client
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
	if exist, _ := c.Store.Get(cont, user.User).Result(); exist != "" {
		return errors.New("already exists")
	}

	if err := c.Store.Set(context.Background(), user.User, user.Password, 20*time.Second).Err(); err != nil {
		return err
	}

	return nil
}

func (c *redisDB) Read(ctx context.Context, user models.User) (string, error) {
	cont, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	val := c.Store.Get(cont, user.User)
	if val.Err() != nil {
		return "", val.Err()
	}
	return val.Val(), nil
}

func (c *redisDB) Update(ctx context.Context, user models.User) error {
	cont, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	if err := c.Store.Get(cont, user.User).Err(); err != nil {
		return errors.New("no such user")
	}
	if err := c.Store.GetSet(context.Background(), user.User, user.Password).Err(); err != nil {
		return err
	}
	return nil
}

func (c *redisDB) Delete(ctx context.Context, user models.User) error {
	cont, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	if err := c.Store.GetDel(cont, user.User).Err(); err != nil {
		return err
	}
	return nil
}
