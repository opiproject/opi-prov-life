package cmd

import (
	"github.com/alknopfler/opi-prov-life/examples/sztp/sztp-agent/pkg/secureAgent"
	"github.com/spf13/cobra"
)

func NewStatusCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Run the status command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return secureAgent.RunCommandStatus()
		},
	}
	return cmd
}
