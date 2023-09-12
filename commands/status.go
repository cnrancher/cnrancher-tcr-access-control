package commands

import (
	"fmt"

	"github.com/cnrancher/tcr-access-control/pkg/config"
	"github.com/cnrancher/tcr-access-control/pkg/tcr"
	"github.com/cnrancher/tcr-access-control/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type statusCmd struct {
	baseCmd
}

func newStatusCmd() *statusCmd {
	cc := &statusCmd{}

	cc.baseCmd.cmd = &cobra.Command{
		Use:   "status",
		Short: "Show status",
		Example: `
    tcr-access-control status`,
		RunE: func(cmd *cobra.Command, args []string) error {
			initializeFlagsConfig(cmd, config.DefaultProvider)
			logrus.Debugf("Debug output enabled")

			err := utils.Init(config.GetString("config"))
			if err != nil {
				return err
			}

			if err = tcr.Init(); err != nil {
				return err
			}
			status, err := tcr.DescribeExternalEndpointStatus()
			if err != nil {
				return err
			}
			if status != nil && status.Response != nil {
				logrus.Printf("External Endpoint Status: %v", utils.Value(status.Response.Status))
				if utils.Value(status.Response.Reason) != "" {
					logrus.Printf("Reason: %v\n", utils.Value(status.Response.Reason))
				}
			}

			securityPolicies, err := tcr.DescribeSecurityPolicies()
			if err != nil {
				return err
			}
			if securityPolicies != nil && securityPolicies.Response != nil &&
				securityPolicies.Response.SecurityPolicySet != nil {
				logrus.Printf("Security Policies:")
				fmt.Printf("INDEX |        CIDR        |  Description\n")
				fmt.Printf("------+--------------------+--------------\n")
				for _, sp := range securityPolicies.Response.SecurityPolicySet {
					fmt.Printf(" %4d | %18s | %v \n",
						utils.Value(sp.PolicyIndex),
						utils.Value(sp.CidrBlock),
						utils.Value(sp.Description))
				}
				fmt.Printf("------+--------------------+--------------\n")
			} else {
				logrus.Infof("No security policies found.")
			}

			return nil
		},
	}

	cc.cmd.Flags().StringP("config", "", utils.TAC_CONFIG_FILE, "Config file")

	return cc
}

func (cc *statusCmd) getCommand() *cobra.Command {
	return cc.cmd
}
