package cmd

import (
	"github.com/alknopfler/opi-prov-life/examples/sztp/sztp-agent/pkg/secureAgent"
	"github.com/spf13/cobra"
)

func NewDisableCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Run the disable command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return secureAgent.RunCommandDisable()
		},
	}
	return cmd
}
