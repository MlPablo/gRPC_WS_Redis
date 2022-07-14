package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/MlPablo/gRPCWebSocket/grpc/crud"
	"github.com/MlPablo/gRPCWebSocket/internal/service"
	"github.com/MlPablo/gRPCWebSocket/internal/store"
)

func main() {
	storage := store.New()
	crud := service.NewCRUDService(storage)
	s := grpc.NewServer()
	pb.RegisterCRUDServer(s, &grpcServer{crud})
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Println(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
