package server

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/vikasbhutra/mcp-server/go-server/internal/agent"
)

type WebsocketConnection struct {
	Conn *websocket.Conn
}

func (c *WebsocketConnection) Send(msg *agent.Message) {
	if err := c.Conn.WriteJSON(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
