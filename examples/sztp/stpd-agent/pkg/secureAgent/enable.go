package secureAgent

import "github.com/alknopfler/opi-prov-life/examples/sztp/stpd-agent/agent"

func RunCommandEnable() *agent.Agent {
	return agent.NewAgent("")
}
