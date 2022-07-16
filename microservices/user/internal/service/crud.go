package service

import (
	"context"

	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/models"
	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/store"
)

type CRUDService interface {
	CreateUser(ctx context.Context, user models.User) error
	UpdateUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, id string) (string, error)
	DeleteUser(ctx context.Context, user string) error
}

type crudService struct {
	storage store.Storage
}

func NewCRUDService(store store.Storage) CRUDService {
	return &crudService{storage: store}
}

func (c *crudService) CreateUser(ctx context.Context, user models.User) error {
	if err := c.storage.NewCRUD().Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (c *crudService) UpdateUser(ctx context.Context, user models.User) error {
	if err := c.storage.NewCRUD().Update(ctx, user); err != nil {
		return err
	}
	return nil
}
func (c *crudService) GetUser(ctx context.Context, id string) (string, error) {
	user, err := c.storage.NewCRUD().Read(ctx, id)
	if err != nil {
		return "", err
	}
	return user, nil
}
func (c *crudService) DeleteUser(ctx context.Context, user string) error {
	if err := c.storage.NewCRUD().Delete(ctx, user); err != nil {
		return err
	}
	return nil
}
