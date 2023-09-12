package commands

import (
	"github.com/spf13/cobra"
)

type baseCmd struct {
	cmd *cobra.Command
}

func (cc *baseCmd) getCommand() *cobra.Command {
	return cc.cmd
}

type cmder interface {
	getCommand() *cobra.Command
}

func addCommands(root *cobra.Command, commands ...cmder) {
	for _, command := range commands {
		cmd := command.getCommand()
		if cmd == nil {
			continue
		}
		root.AddCommand(cmd)
	}
}
