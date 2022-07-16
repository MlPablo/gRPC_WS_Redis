package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	pb "github.com/MlPablo/gRPCWebSocket/microservices/user/grpc/crud"
	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/models"
	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/service"
)

type GrpcServer struct {
	S service.CRUDService
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

func (g *GrpcServer) CreateUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	user, err := UnmarshalRequest(req)
	//log.Println(user)
	if err != nil {
		return &pb.Response{Success: false}, err
	}
	if len(user.User) == 0 || len(user.Password) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	if err := g.S.CreateUser(ctx, user); err != nil {
		log.Println(user)
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true}, nil
}

func (g *GrpcServer) UpdateUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	user, err := UnmarshalRequest(req)
	if err != nil {
		return &pb.Response{Success: false}, err
	}
	if len(user.User) == 0 || len(user.Password) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	if err := g.S.UpdateUser(ctx, user); err != nil {
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true}, nil
}
func (g *GrpcServer) GetUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	user, err := UnmarshalRequest(req)
	if err != nil {
		return &pb.Response{Success: false}, err
	}
	if len(user.User) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	userPassword, err := g.S.GetUser(ctx, user.User)
	if err != nil {
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true, Body: userPassword}, nil
}
func (g *GrpcServer) DeleteUser(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	user, err := UnmarshalRequest(req)
	if err != nil {
		return &pb.Response{Success: false}, err
	}
	if len(user.User) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	if err := g.S.DeleteUser(ctx, user.User); err != nil {
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true}, nil
}
