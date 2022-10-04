package secureAgent

import "github.com/alknopfler/opi-prov-life/examples/sztp/stpd-agent/agent"

func RunCommandStatus() *agent.Agent {
	return agent.NewAgent("")
}
