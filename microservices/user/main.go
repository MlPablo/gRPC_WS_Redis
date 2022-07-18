package main

import (
	"log"
	"net"
	"os"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/MlPablo/gRPCWebSocket/microservices/user/api"
	pb "github.com/MlPablo/gRPCWebSocket/microservices/user/grpc/crud"
	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/service"
	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/store"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	session, err := NewScyllaSession()
	if err != nil {
		log.Fatal(err)
	}
	storage := store.NewScylla(session)
	crud := service.NewCRUDService(storage)
	serv := api.GrpcServer{S: crud}
	s := grpc.NewServer()
	pb.RegisterCRUDServer(s, &serv)
	l, err := net.Listen("tcp", os.Getenv("GRPC_USER_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func NewScyllaSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster(os.Getenv("SCYLLA_URL"))
	cluster.ProtoVersion = 4
	//cluster.Keyspace = "myapp"
	cluster.NumConns = 3
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	err = session.Query("CREATE KEYSPACE IF NOT EXISTS myapp " +
		"WITH REPLICATION = {" +
		"'class' : 'NetworkTopologyStrategy'," +
		"'replication_factor': 3};").Exec()
	if err != nil {
		return nil, err
	}
	err = session.Query("CREATE TABLE IF NOT EXISTS myapp.users " +
		"(name text, password text, register_time timestamp, PRIMARY KEY (name));").Exec()
	if err != nil {
		return nil, err
	}
	return session, nil
}
