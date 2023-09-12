package commands

import (
	"os"

	"github.com/cnrancher/tcr-access-control/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Execute(args []string) {
	tacCmd := newTACCmd()
	tacCmd.addCommands()
	tacCmd.cmd.SetArgs(args)

	_, err := tacCmd.cmd.ExecuteC()
	if err != nil {
		os.Exit(1)
	}
}

type tacCmd struct {
	baseCmd
}

func newTACCmd() *tacCmd {
	cc := &tacCmd{}

	cc.baseCmd.cmd = &cobra.Command{
		Use:   "tcr-access-control",
		Short: "tcr-access-control usage",
		Long: `tcr-access-control is a tool for manage the
Tencent Cloud TCR access security policies.

https://github.com/cnrancher/tcr-access-control`,
		Run: func(cmd *cobra.Command, args []string) {
			initializeFlagsConfig(cmd, config.DefaultProvider)
			logrus.Debugf("Debug output enabled")
			// show help message only
			cmd.Help()
		},
	}
	cc.cmd.CompletionOptions = cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	}
	cc.cmd.Version = getVersion()
	cc.cmd.SilenceUsage = true

	cc.cmd.PersistentFlags().BoolP("debug", "", false, "enable debug output")

	return cc
}

func (cc *tacCmd) getCommand() *cobra.Command {
	return cc.cmd
}

func (cc *tacCmd) addCommands() {
	addCommands(
		cc.cmd,
		newAllowCmd(),
		newRemoveCmd(),
		newStatusCmd(),
		newConfigCmd(),
		newVersionCmd(),
	)
}
