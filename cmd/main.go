package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrader websocket Upgrade
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
