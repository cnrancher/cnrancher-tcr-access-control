package commands

import (
	"github.com/cnrancher/tcr-access-control/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type deleteCmd struct {
	baseCmd
}

func newDeleteCmd() *deleteCmd {
	cc := &deleteCmd{}

	cc.baseCmd.cmd = &cobra.Command{
		Use:     "delete",
		Short:   "Remove IP address from security policy",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			initializeFlagsConfig(cmd, config.DefaultProvider)
			logrus.Debugf("Debug output enabled")

			return nil
		},
	}

	return cc
}

func (cc *deleteCmd) getCommand() *cobra.Command {
	return cc.cmd
}
