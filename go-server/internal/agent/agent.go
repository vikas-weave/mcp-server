package agent

type Connection interface {
	Send(msg *Message)
}

type Agent struct {
	ID           string            `json:"id"`
	Capabilities []string          `json:"capabilities"`
	Metadata     map[string]string `json:"metadata"`
	Conn         Connection
}

var agentRegistry = make(map[string]*Agent)

func RegisterAgent(agent *Agent) {
	agentRegistry[agent.ID] = agent
}

func GetAgent(id string) (*Agent, bool) {
	agent, exists := agentRegistry[id]
	return agent, exists
}
