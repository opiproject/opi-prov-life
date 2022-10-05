package cmd

import (
	"github.com/alknopfler/opi-prov-life/examples/sztp/sztp-agent/pkg/secureAgent"
	"github.com/spf13/cobra"
)

func NewDaemonCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "daemon",
		Short: "Run the daemon command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return secureAgent.RunCommandDaemon()
		},
	}
	return cmd
}
