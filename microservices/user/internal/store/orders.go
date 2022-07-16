package store

import (
	"context"
	"errors"
	"time"

	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/models"
)

type Orders interface {
	Create(ctx context.Context, user models.Order) error
}

type orders struct {
	*storage
}

func (o *orders) Create(ctx context.Context, order models.Order) error {
	cont, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	if exist, _ := o.store.Get(cont, order.Name).Result(); exist != "" {
		return errors.New("already exists")
	}

	if err := o.store.Set(cont, order.Name, order.Type, 0).Err(); err != nil {
		return err
	}
	return nil
}
