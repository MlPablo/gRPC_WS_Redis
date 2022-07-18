package service

import (
	"context"

	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/models"
	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/store"
)

type CRUDService interface {
	CreateUser(ctx context.Context, user models.User) error
	UpdateUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, user models.User) (string, error)
	DeleteUser(ctx context.Context, user models.User) error
}

type crudService struct {
	auth    *authorization
	storage store.Storage
}

func NewCRUDService(store store.Storage) CRUDService {
	return &crudService{auth: NewAuthorization(), storage: store}
}

func (c *crudService) CreateUser(ctx context.Context, user models.User) error {
	if err := c.storage.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (c *crudService) UpdateUser(ctx context.Context, user models.User) error {
	if err := c.storage.Update(ctx, user); err != nil {
		return err
	}
	return nil
}
func (c *crudService) GetUser(ctx context.Context, user models.User) (string, error) {
	c.auth.Authorize(ctx, user.User)
	pass, err := c.storage.Read(ctx, user)
	if err != nil {
		return "", err
	}
	return pass, nil
}
func (c *crudService) DeleteUser(ctx context.Context, user models.User) error {
	if err := c.storage.Delete(ctx, user); err != nil {
		return err
	}
	return nil
}
