package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/MlPablo/gRPCWebSocket/cmd/handler"
)

// SetUpRoutes define endpoints
func SetUpRoutes() {
	http.HandleFunc("/", handler.ConnectionChecker)
}

// run server on port 2828
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	SetUpRoutes()
	log.Fatal(http.ListenAndServe(os.Getenv("WEBSOCKET_URL"), nil))
}
