package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/MlPablo/gRPCWebSocket/cmd/grpc/service"
)

// upgrader websocket Upgrade
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Code struct {
	OpCode int `form:"opcode" json:"opcode" binding:"required" validate:"required"`
}

// ConnectionChecker handler for / endpoint with upgrading it to websocket and listening for messages
func ConnectionChecker(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client connected ...")
	gRPC, err := grpc.Dial(os.Getenv("GRPC_SERVICE_HOST")+os.Getenv("GRPC_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewRegisterClient(gRPC)
	reader(ws, client)
}

// reader handler for getting and sending messages by socket.
func reader(conn *websocket.Conn, client pb.RegisterClient) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			continue
		}
		code := &Code{}
		if err := json.Unmarshal(msg, &code); err != nil {
			log.Println(err)
			continue
		}
		res, _ := client.Get(context.Background(), &pb.Request{Code: int32(code.OpCode), Body: msg})
		if err != nil {
			conn.WriteMessage(1, []byte(err.Error()))
			log.Println(err)
		} else {
			answer, _ := json.Marshal(fmt.Sprintf("Succes: %t | Body: %s | Error: %s", res.GetSuccess(), res.GetBody(), res.GetError()))
			conn.WriteMessage(1, answer)
		}
	}
}
