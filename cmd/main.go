package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/MlPablo/gRPCWebSocket/grpc/crud"
)

// PingPong struct for ping pong message
type Request struct {
	OpCode int        `form:"opcode" json:"opcode" binding:"required" validate:"required"`
	Body   pb.Request `form:"body" json:"body" binding:"required" validate:"required"`
}

// upgrader websocket Upgrade
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleResponse(res *pb.Response, conn *websocket.Conn, err error) {
	if err != nil {
		js, _ := json.Marshal(map[string]string{"error": err.Error()})
		conn.WriteMessage(1, js)
	} else {
		js, _ := json.Marshal(&res)
		conn.WriteMessage(1, js)
	}
}

// reader handler for getting and sending messages by socket.
func reader(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		req := &Request{}
		if err := json.Unmarshal(msg, &req); err != nil {
			log.Fatal(err)
			return
		}

		gRPC, err := grpc.Dial(":80", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
			return
		}
		client := pb.NewCRUDClient(gRPC)
		switch req.OpCode {
		case 1:
			res, err := client.CreateUser(context.Background(), &req.Body)
			handleResponse(res, conn, err)
			log.Println(res, err)
		case 2:
			res, err := client.UpdateUser(context.Background(), &req.Body)
			handleResponse(res, conn, err)
			log.Println(res, err)
		case 3:
			res, err := client.GetUser(context.Background(), &req.Body)
			handleResponse(res, conn, err)
			log.Println(res, err)
		case 4:
			res, err := client.DeleteUser(context.Background(), &req.Body)
			handleResponse(res, conn, err)
			log.Println(res, err)
		default:
			js, _ := json.Marshal(map[string]string{"error": "unknown OpCode"})
			conn.WriteMessage(1, js)
		}
	}
}

// ConnectionChecker handler for /ws endpoint with upgrading it to websocket and listening for messages
func ConnectionChecker(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client connected ...")
	reader(ws)
}

// SetUpRoutes define endpoints
func SetUpRoutes() {
	http.HandleFunc("/", ConnectionChecker)
}

// run server on port 28000
func main() {
	SetUpRoutes()
	log.Fatal(http.ListenAndServe(":28000", nil))
}
