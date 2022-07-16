package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/MlPablo/gRPCWebSocket/microservices/entry/grpc/service"
	service2 "github.com/MlPablo/gRPCWebSocket/microservices/entry/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	gRPCUser, err := grpc.Dial(os.Getenv("GRPC_USER_HOST")+os.Getenv("GRPC_USER_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	gRPCOrder, err := grpc.Dial(os.Getenv("GRPC_ORDER_HOST")+os.Getenv("GRPC_ORDER_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	service := service2.NewUserClient(gRPCUser, gRPCOrder)
	s := grpc.NewServer()
	pb.RegisterRegisterServer(s, service)
	l, err := net.Listen("tcp", os.Getenv("GRPC_SERVICE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
