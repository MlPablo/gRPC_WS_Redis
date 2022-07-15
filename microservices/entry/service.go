package main

import (
	"context"

	"google.golang.org/grpc"

	pb2 "github.com/MlPablo/gRPCWebSocket/microservices/entry/grpc/crud"
	pb1 "github.com/MlPablo/gRPCWebSocket/microservices/entry/grpc/order"
	pb "github.com/MlPablo/gRPCWebSocket/microservices/entry/grpc/service"
)

type UserClient struct {
	CrudClient  pb2.CRUDClient
	OrderClient pb1.OrderClient
}

func NewUserClient(cc grpc.ClientConnInterface, oc grpc.ClientConnInterface) *UserClient {
	return &UserClient{
		CrudClient:  pb2.NewCRUDClient(cc),
		OrderClient: pb1.NewOrderClient(oc),
	}
}

func (u *UserClient) Get(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	switch req.GetCode() {
	case 1:
		res, err := u.CrudClient.CreateUser(ctx, &pb2.Request{Body: req.Body})
		return &pb.Response{Success: res.GetSuccess(), Body: res.GetBody()}, err
	case 2:
		res, err := u.CrudClient.GetUser(ctx, &pb2.Request{Body: req.Body})
		return &pb.Response{Success: res.GetSuccess(), Body: res.GetBody()}, err
	case 3:
		res, err := u.CrudClient.UpdateUser(ctx, &pb2.Request{Body: req.Body})
		return &pb.Response{Success: res.GetSuccess(), Body: res.GetBody()}, err
	case 4:
		res, err := u.CrudClient.DeleteUser(ctx, &pb2.Request{Body: req.Body})
		return &pb.Response{Success: res.GetSuccess(), Body: res.GetBody()}, err
	case 5:
		res, err := u.OrderClient.CreateOrder(ctx, &pb1.Request{Body: req.Body})
		return &pb.Response{Success: res.GetSuccess()}, err
	default:
		return &pb.Response{}, nil
	}
	//return &pb.Response{}, nil
}
