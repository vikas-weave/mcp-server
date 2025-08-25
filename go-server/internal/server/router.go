package server

import (
	"log"

	"github.com/vikasbhutra/mcp-server/go-server/internal/agent"
)

func RouteMessage(msg *agent.Message) {
	targetAgent, exists := agent.GetAgent(msg.To)
	if !exists {
		log.Printf("Target agent not found: %s", msg.To)
		return
	}
	targetAgent.Conn.Send(msg)
}
