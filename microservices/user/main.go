package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/MlPablo/gRPCWebSocket/microservices/user/api"
	pb "github.com/MlPablo/gRPCWebSocket/microservices/user/grpc/crud"
	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/service"
	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/store"
)

func main() {
	storage := store.New()
	crud := service.NewCRUDService(storage)
	serv := api.GrpcServer{S: crud}
	s := grpc.NewServer()
	pb.RegisterCRUDServer(s, &serv)
	l, err := net.Listen("tcp", ":81")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
