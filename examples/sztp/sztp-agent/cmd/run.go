package cmd

import (
	"github.com/alknopfler/opi-prov-life/examples/sztp/sztp-agent/pkg/secureAgent"
	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Exec the run command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return secureAgent.RunCommandRun()
		},
	}
	return cmd
}
