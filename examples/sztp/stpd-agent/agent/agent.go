package agent

var (
	AgentConfig Agent
)

type Agent struct {
	Error     error
	ConfigURL string
	//TBD the rest of fields
}

func NewAgent(configURL string) *Agent {
	return &Agent{ConfigURL: configURL, Error: nil}
}

func (a Agent) GetConfigURL() string {
	return a.ConfigURL
}

func (a Agent) GetError() error {
	return a.Error
}

func (a Agent) GetErrorDesc() string {
	return a.Error.Error()
}
