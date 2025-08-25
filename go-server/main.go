package main

import (
	"log"
	"net/http"

	"github.com/vikasbhutra/mcp-server/go-server/internal/server"
)

func main() {
	http.HandleFunc("/ws", server.ServeHTTP)
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
