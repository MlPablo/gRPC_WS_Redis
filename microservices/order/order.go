package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	pb "github.com/MlPablo/gRPCWebSocket/microservices/order/grpc/order"

	"github.com/MlPablo/gRPCWebSocket/internal/models"
	"github.com/MlPablo/gRPCWebSocket/internal/service"
)

type grpcServer struct {
	s service.OrderService
}

type Request struct {
	Body models.Order `json:"body"`
}

func UnmarshalRequest(req *pb.Request) (models.Order, error) {
	reqBody := Request{}
	if err := json.Unmarshal(req.Body, &reqBody); err != nil {
		return reqBody.Body, err
	}
	return reqBody.Body, nil
}

func (g *grpcServer) CreateOrder(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	order, err := UnmarshalRequest(req)
	if err != nil {
		return &pb.Response{}, err
	}
	if len(order.Type) == 0 || len(order.Name) == 0 {
		return &pb.Response{Success: false}, errors.New("unable to validate data")
	}
	log.Println(req)
	if err := g.s.AddOrder(ctx, order); err != nil {
		return &pb.Response{Success: false}, err
	}
	return &pb.Response{Success: true}, nil
}
