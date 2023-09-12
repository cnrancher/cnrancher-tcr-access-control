package commands

import (
	"fmt"
	"net"

	"github.com/cnrancher/tcr-access-control/pkg/config"
	"github.com/cnrancher/tcr-access-control/pkg/tcr"
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
		Short: "Add one IPv4 address (CIDR block) to security policy",
		Example: `  tcr-access-control allow \
	--ip="192.168.0.0/24" \
	--description="Example"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			initializeFlagsConfig(cmd, config.DefaultProvider)
			logrus.Debugf("Debug output enabled")
			cidr := config.GetString("ip")
			if cidr == "" {
				logrus.Errorf("ip not provided")
				cc.cmd.Usage()
				return fmt.Errorf("ip not provided")
			}

			var err error
			// Input IP should be a valid IPv4 address or CIDR block
			ip := net.ParseIP(cidr)
			if ip == nil {
				ip, _, err = net.ParseCIDR(cidr)
				if err != nil {
					return fmt.Errorf("invalid format: %w", err)
				}
			}
			if ip.To4() == nil {
				return fmt.Errorf(
					"invalid IP %q, only IPv4 allowed", cidr)
			}

			if err := utils.Init(config.GetString("config")); err != nil {
				return err
			}
			if err := tcr.Init(); err != nil {
				return err
			}
			response, err := tcr.CreateSecurityPolicy(
				cidr, config.GetString("description"))
			if err != nil {
				return fmt.Errorf(
					"CreateMultipleSecurityPolicy failed: %w", err)
			}
			logrus.Debugf("%v", response.ToJsonString())
			logrus.Infof("Successfully add %q to security policy", cidr)

			return nil
		},
	}

	cc.cmd.Flags().StringP("config", "", utils.TAC_CONFIG_FILE, "Config file")
	cc.cmd.Flags().StringP("ip", "", "",
		"IPv4 address (CIDR block) (required)")
	cc.cmd.Flags().StringP("description", "d", "", "Description")

	return cc
}

func (cc *allowCmd) getCommand() *cobra.Command {
	return cc.cmd
}
