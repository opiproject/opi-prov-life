package cmd

import (
	"github.com/alknopfler/opi-prov-life/examples/sztp/sztp-agent/pkg/secureAgent"
	"github.com/spf13/cobra"
)

func NewEnableCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Run the enable command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return secureAgent.RunCommandEnable()
		},
	}
	return cmd
}
