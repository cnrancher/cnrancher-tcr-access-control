package commands

import (
	"fmt"

	"github.com/cnrancher/tcr-access-control/pkg/config"
	"github.com/cnrancher/tcr-access-control/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type versionCmd struct {
	baseCmd
}

func newVersionCmd() *versionCmd {
	cc := &versionCmd{}

	cc.baseCmd.cmd = &cobra.Command{
		Use:     "version",
		Short:   "Show version",
		Example: "  tcr-access-control version",
		Run: func(cmd *cobra.Command, args []string) {
			initializeFlagsConfig(cmd, config.DefaultProvider)
			logrus.Debugf("Debug output enabled")
			fmt.Printf("tcr-access-control version %s\n", getVersion())
		},
	}

	return cc
}

func (cc *versionCmd) getCommand() *cobra.Command {
	return cc.cmd
}

func getVersion() string {
	if utils.GitCommit != "" {
		return fmt.Sprintf("%s - %s", utils.Version, utils.GitCommit)
	}
	return utils.Version
}
