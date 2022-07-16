package main

import (
	"log"
	"net/http"

	"github.com/MlPablo/gRPCWebSocket/cmd/handler"
)

// SetUpRoutes define endpoints
func SetUpRoutes() {
	http.HandleFunc("/", handler.ConnectionChecker)
}

// run server on port 28000
func main() {
	SetUpRoutes()
	log.Fatal(http.ListenAndServe(":2828", nil))
}
