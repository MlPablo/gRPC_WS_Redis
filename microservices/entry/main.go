package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/MlPablo/gRPCWebSocket/microservices/entry/grpc/service"
)

func main() {
	gRPC, err := grpc.Dial(":81", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	gRPC1, err := grpc.Dial(":82", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	service := NewUserClient(gRPC, gRPC1)
	s := grpc.NewServer()
	pb.RegisterRegisterServer(s, service)
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
