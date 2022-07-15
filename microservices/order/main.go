package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/MlPablo/gRPCWebSocket/internal/service"
	"github.com/MlPablo/gRPCWebSocket/internal/store"
	pb "github.com/MlPablo/gRPCWebSocket/microservices/order/grpc/order"
)

func main() {
	storage := store.New()
	order := service.NewOrderService(storage)
	serv := grpcServer{order}
	s := grpc.NewServer()
	pb.RegisterOrderServer(s, &serv)
	l, err := net.Listen("tcp", ":82")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
