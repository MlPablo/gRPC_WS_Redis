package service

import (
	"context"

	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/models"
	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/store"
)

type authorization struct {
	redisDB store.Storage
}

func NewAuthorization() *authorization {
	return &authorization{store.NewRedis()}
}

func (a *authorization) Authorize(ctx context.Context, user models.User) (string, bool) {
	if string, err := a.redisDB.Read(ctx, user); err == nil {
		return string, true
	}
	return "", false
}

func (a *authorization) AddClient(ctx context.Context, user models.User) error {
	return a.redisDB.Create(ctx, user)
}
