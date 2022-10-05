package secureAgent

type Agent struct {
	ConfigURL string
	//TBD the rest of fields
}

func NewAgent(configURL string) *Agent {
	return &Agent{ConfigURL: configURL}
}

func (a Agent) GetConfigURL() string {
	return a.ConfigURL
}
