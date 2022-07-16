package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/MlPablo/gRPCWebSocket/microservices/entry/grpc/service"
	service2 "github.com/MlPablo/gRPCWebSocket/microservices/entry/service"
)

func main() {
	gRPC, err := grpc.Dial("dns:///ws_grpc-user-1:81", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	gRPC1, err := grpc.Dial("dns:///ws_grpc-order-1:82", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	service := service2.NewUserClient(gRPC, gRPC1)
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
