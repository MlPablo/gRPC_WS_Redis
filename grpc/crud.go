package main

import (
	"context"
	"errors"

	pb "github.com/MlPablo/gRPCWebSocket/grpc/crud"
	"github.com/MlPablo/gRPCWebSocket/internal/models"
	"github.com/MlPablo/gRPCWebSocket/internal/service"
)

type grpcServer struct {
	s service.CRUDService
}

func (g *grpcServer) CreateUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	if len(req.Name) == 0 || len(req.Password) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	user := models.User{User: req.Name, Password: req.Password}
	if err := g.s.CreateUser(ctx, user); err != nil {
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true}, nil
}

func (g *grpcServer) UpdateUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	if len(req.Name) == 0 || len(req.Password) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	user := models.User{User: req.Name, Password: req.Password}
	if err := g.s.UpdateUser(ctx, user); err != nil {
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true}, nil
}
func (g *grpcServer) GetUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	if len(req.Name) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	user, err := g.s.GetUser(ctx, req.GetName())
	if err != nil {
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true, Body: &pb.Request{Password: user}}, nil
}
func (g *grpcServer) DeleteUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	if len(req.Name) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	if err := g.s.DeleteUser(ctx, req.GetName()); err != nil {
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true}, nil
}

func (g *grpcServer) mustEmbedUnimplementedCRUDServer() {
	panic("implement me")
}
