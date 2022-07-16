package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/MlPablo/gRPCWebSocket/microservices/order/internal/service"
	"github.com/MlPablo/gRPCWebSocket/microservices/order/internal/store"

	"github.com/MlPablo/gRPCWebSocket/microservices/order/api"
	pb "github.com/MlPablo/gRPCWebSocket/microservices/order/grpc/order"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	storage := store.New()
	order := service.NewOrderService(storage)
	serv := api.GrpcServer{S: order}
	s := grpc.NewServer()
	pb.RegisterOrderServer(s, &serv)
	l, err := net.Listen("tcp", os.Getenv("GRPC_ORDER_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
