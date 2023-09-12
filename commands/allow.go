package commands

import (
	"fmt"

	"github.com/cnrancher/tcr-access-control/pkg/config"
	"github.com/cnrancher/tcr-access-control/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type allowCmd struct {
	baseCmd
}

func newAllowCmd() *allowCmd {
	cc := &allowCmd{}

	cc.baseCmd.cmd = &cobra.Command{
		Use:   "allow",
		Short: "Add IP address (CIDR block) to security policy",
		Example: `
	tcr-access-control allow \
		--ip 192.168.0.0/24,10.0.0.0/8,172.16.0.0/16 \
		--description="Example"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			initializeFlagsConfig(cmd, config.DefaultProvider)
			logrus.Debugf("Debug output enabled")

			err := utils.Init(config.GetString("config"))
			if err != nil {
				return err
			}

			if config.GetString("ip") == "" {
				logrus.Errorf("ip not provided")
				cc.cmd.Usage()
				return fmt.Errorf("ip not provided")
			}

			return nil
		},
	}

	cc.cmd.Flags().StringP("config", "", utils.TAC_CONFIG_FILE, "Config file")
	cc.cmd.Flags().StringP("ip", "i", "", "IP address (CIDR block), split by comma")
	cc.cmd.Flags().StringP("description", "d", "", "Description")

	return cc
}

func (cc *allowCmd) getCommand() *cobra.Command {
	return cc.cmd
}
