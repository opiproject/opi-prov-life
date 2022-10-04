package main

import (
	"github.com/TwiN/go-color"
	"github.com/alknopfler/opi-prov-life/examples/sztp/stpd-agent/cmd"

	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	command := newCommand()
	if err := command.Execute(); err != nil {
		log.Fatalf(color.InRed("[ERROR]")+"%s", err.Error())
	}
}

func newCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "sztp",
		Short: "aztp is the agent command line interface to work with the sztp workflow",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	c.AddCommand(cmd.NewStatusCommand())
	c.AddCommand(cmd.NewEnableCommand())
	c.AddCommand(cmd.NewDisableCommand())

	return c
}
