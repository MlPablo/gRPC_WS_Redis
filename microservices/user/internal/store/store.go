package store

import (
	"context"

	"github.com/gocql/gocql"

	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/models"
	redisDB "github.com/MlPablo/gRPCWebSocket/microservices/user/internal/store/redis"
	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/store/scylla"
)

type Storage interface {
	Create(ctx context.Context, user models.User) error
	Read(ctx context.Context, user models.User) (string, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, user models.User) error
}

func NewRedis() Storage {
	return redisDB.New()
}

func NewScylla(session *gocql.Session) Storage {
	return scylla.New(session)
}
