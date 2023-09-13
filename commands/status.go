package commands

import (
	"fmt"

	"github.com/cnrancher/tcr-access-control/pkg/cmdconfig"
	"github.com/cnrancher/tcr-access-control/pkg/policystatus"
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
		Use:     "status",
		Short:   "Show status",
		Example: `  tcr-access-control status`,
		RunE: func(cmd *cobra.Command, args []string) error {
			initializeFlagsConfig(cmd, cmdconfig.DefaultProvider)
			logrus.Debugf("Debug output enabled")

			err := utils.Init(cmdconfig.GetString("config"))
			if err != nil {
				return fmt.Errorf("utils Init failed: %w", err)
			}

			if err = tcr.Init(); err != nil {
				return fmt.Errorf("TCR Init failed: %w", err)
			}

			statusRes, err := tcr.DescribeExternalEndpointStatus()
			if err != nil {
				return fmt.Errorf("DescribeExternalEndpointStatus: %w", err)
			}
			logrus.Debugf("%v", statusRes.ToJsonString())

			status := policystatus.Status{}
			if statusRes != nil && statusRes.Response != nil {
				status.Status = utils.Value(statusRes.Response.Status)
				status.Reason = utils.Value(statusRes.Response.Reason)
			}

			securityPolicies, err := tcr.DescribeSecurityPolicies()
			if err != nil {
				// DescribeSecurityPolicies will fail when External Endpoint
				// status is not Open
				logrus.Warnf("Error when DescribeSecurityPolicies: %v", err)
				logrus.Warnf(
					"Please ensure the External Endpoint Status is 'Opened'")
				err = nil
			}

			if securityPolicies != nil && securityPolicies.Response != nil &&
				securityPolicies.Response.SecurityPolicySet != nil {
				for _, s := range securityPolicies.Response.SecurityPolicySet {
					p := policystatus.SecurityPolicy{
						Index:       utils.Value(s.PolicyIndex),
						CIDR:        utils.Value(s.CidrBlock),
						Description: utils.Value(s.Description),
					}
					status.Policies = append(status.Policies, p)
				}
			}

			if cmdconfig.GetBool("json") {
				fmt.Println(status.Json())
			} else {
				fmt.Print(status.String())
			}

			return nil
		},
	}

	cc.cmd.Flags().StringP("config", "", utils.TAC_CONFIG_FILE, "Config file")
	cc.cmd.Flags().BoolP("json", "", false, "Json format")

	return cc
}

func (cc *statusCmd) getCommand() *cobra.Command {
	return cc.cmd
}

func (cc *statusCmd) printStatusMsg() {

}
