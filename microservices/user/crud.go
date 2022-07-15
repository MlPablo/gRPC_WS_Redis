package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/MlPablo/gRPCWebSocket/internal/models"
	"github.com/MlPablo/gRPCWebSocket/internal/service"
	pb "github.com/MlPablo/gRPCWebSocket/microservices/user/grpc/crud"
)

type grpcServer struct {
	s service.CRUDService
}

type Request struct {
	Body models.User `json:"body"`
}

func UnmarshalRequest(req *pb.Request) (models.User, error) {
	reqBody := Request{}
	if err := json.Unmarshal(req.Body, &reqBody); err != nil {
		return reqBody.Body, err
	}
	//log.Println(reqBody.Body)
	return reqBody.Body, nil
}

func (g *grpcServer) CreateUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	user, err := UnmarshalRequest(req)
	//log.Println(user)
	if err != nil {
		return &pb.Response{Success: false}, err
	}
	if len(user.User) == 0 || len(user.Password) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	if err := g.s.CreateUser(ctx, user); err != nil {
		log.Println(user)
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true}, nil
}

func (g *grpcServer) UpdateUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	user, err := UnmarshalRequest(req)
	if err != nil {
		return &pb.Response{Success: false}, err
	}
	if len(user.User) == 0 || len(user.Password) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	if err := g.s.UpdateUser(ctx, user); err != nil {
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true}, nil
}
func (g *grpcServer) GetUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	user, err := UnmarshalRequest(req)
	if err != nil {
		return &pb.Response{Success: false}, err
	}
	if len(user.User) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	userPassword, err := g.s.GetUser(ctx, user.User)
	if err != nil {
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true, Body: userPassword}, nil
}
func (g *grpcServer) DeleteUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	user, err := UnmarshalRequest(req)
	if err != nil {
		return &pb.Response{Success: false}, err
	}
	if len(user.User) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	if err := g.s.DeleteUser(ctx, user.User); err != nil {
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true}, nil
}
