package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/vikas-weave/mcp-server/go-server/internal/agent"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	defer conn.Close()

	client := &WebsocketConnection{Conn: conn}

	// Read the first message to get the agent's ID
	var initialMsg agent.Message
	if err := conn.ReadJSON(&initialMsg); err != nil {
		log.Printf("Error reading initial message: %v", err)
		return
	}

	// Register the agent
	newAgent := &agent.Agent{
		ID:   initialMsg.From,
		Conn: client,
	}
	agent.RegisterAgent(newAgent)

	for {
		var msg agent.Message
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		RouteMessage(&msg)
	}
}
