package commands

import (
	"encoding/json"
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
		Use:     "status",
		Short:   "Show status",
		Example: `  tcr-access-control status`,
		RunE: func(cmd *cobra.Command, args []string) error {
			initializeFlagsConfig(cmd, config.DefaultProvider)
			logrus.Debugf("Debug output enabled")

			err := utils.Init(config.GetString("config"))
			if err != nil {
				return fmt.Errorf("utils Init failed: %w", err)
			}

			if err = tcr.Init(); err != nil {
				return fmt.Errorf("TCR Init failed: %w", err)
			}
			status, err := tcr.DescribeExternalEndpointStatus()
			if err != nil {
				return fmt.Errorf("DescribeExternalEndpointStatus: %w", err)
			}
			if status != nil && status.Response != nil {
				if !config.GetBool("json") {
					logrus.Printf("External Endpoint Status: %v",
						utils.Value(status.Response.Status))
					if utils.Value(status.Response.Reason) != "" {
						logrus.Printf("Reason: %v\n",
							utils.Value(status.Response.Reason))
					}
				}
			}

			securityPolicies, err := tcr.DescribeSecurityPolicies()
			if err != nil {
				return fmt.Errorf("DescribeSecurityPolicies: %w", err)
			}
			if securityPolicies != nil && securityPolicies.Response != nil &&
				securityPolicies.Response.SecurityPolicySet != nil {
				set := securityPolicies.Response.SecurityPolicySet
				if config.GetBool("json") {
					type policy struct {
						Index       int64  `json:"index"`
						CIDR        string `json:"cidr"`
						Description string `json:"description"`
					}

					policies := []policy{}
					for _, s := range set {
						policies = append(policies, policy{
							Index:       utils.Value(s.PolicyIndex),
							CIDR:        utils.Value(s.CidrBlock),
							Description: utils.Value(s.Description),
						})
					}
					data, _ := json.MarshalIndent(policies, "", "    ")
					fmt.Printf("%v\n", string(data))
				} else {
					logrus.Printf("Security Policies:")
					fmt.Printf("INDEX |        CIDR        |  Description\n")
					fmt.Printf("------+--------------------+--------------\n")
					for _, s := range set {
						fmt.Printf(" %4d | %18s | %v \n",
							utils.Value(s.PolicyIndex),
							utils.Value(s.CidrBlock),
							utils.Value(s.Description),
						)
					}
					fmt.Printf("------+--------------------+--------------\n")
				}
			} else {
				return fmt.Errorf("No security policies found.")
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
