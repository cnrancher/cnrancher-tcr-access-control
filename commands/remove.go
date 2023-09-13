package commands

import (
	"fmt"
	"net"

	"github.com/cnrancher/tcr-access-control/pkg/cmdconfig"
	"github.com/cnrancher/tcr-access-control/pkg/tcr"
	"github.com/cnrancher/tcr-access-control/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type removeCmd struct {
	baseCmd
}

func newRemoveCmd() *removeCmd {
	cc := &removeCmd{}

	cc.baseCmd.cmd = &cobra.Command{
		Use:     "remove",
		Short:   "Remove one IPv4 address from security policy",
		Example: `  tcr-access-control remove --ip="192.168.0.0/24"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			initializeFlagsConfig(cmd, cmdconfig.DefaultProvider)
			logrus.Debugf("Debug output enabled")

			policy := cmdconfig.GetString("ip")
			if policy == "" {
				logrus.Errorf("ip (CIDR block) not provided")
				cc.cmd.Usage()
				return fmt.Errorf("ip not provided")
			}
			index := cmdconfig.GetInt64("index")
			if cmdconfig.GetInt64("index") < 0 {
				logrus.Errorf("index not provided, " +
					"use 'tcr-access-control show' to get the index of policy")
				cc.cmd.Usage()
				return fmt.Errorf("index not provided")
			}

			var err error
			// Input IP should be a valid IPv4 address or CIDR block
			ip := net.ParseIP(policy)
			if ip == nil {
				ip, _, err = net.ParseCIDR(policy)
				if err != nil {
					return fmt.Errorf("invalid format: %w", err)
				}
			}
			if ip.To4() == nil {
				return fmt.Errorf(
					"invalid IP %q, only IPv4 allowed", policy)
			}

			logrus.Debugf("IP: %v", policy)

			if err := utils.Init(cmdconfig.GetString("config")); err != nil {
				return err
			}
			if err := tcr.Init(); err != nil {
				return err
			}

			var version string
			policiesRes, err := tcr.DescribeSecurityPolicies()
			if err != nil {
				return fmt.Errorf("DescribeSecurityPolicies failed: %v", err)
			}
			if policiesRes == nil || policiesRes.Response == nil ||
				policiesRes.Response.SecurityPolicySet == nil ||
				len(policiesRes.Response.SecurityPolicySet) == 0 {
				logrus.Errorf("Failed to query security policies")
				logrus.Errorf(
					"Please ensure the External Endpoint status is Opened")
				return fmt.Errorf("No existing security policies from server")
			}

			for _, p := range policiesRes.Response.SecurityPolicySet {
				if utils.Value(p.CidrBlock) != policy {
					continue
				}
				if index != utils.Value(p.PolicyIndex) {
					continue
				}
				logrus.Debugf("Find security policy index %v, %q, version: %q",
					utils.Value(p.PolicyIndex),
					utils.Value(p.CidrBlock),
					utils.Value(p.PolicyVersion))
				version = utils.Value(p.PolicyVersion)
			}
			if version == "" {
				logrus.Errorf("Failed to find existing security policy "+
					"index [%v] CIDR [%v] from server",
					index, policy)
				logrus.Errorf("Use 'tcr-access-control status' to find " +
					"the existing security policy to delete")
				return fmt.Errorf("security policy not found")
			}

			if !cmdconfig.GetBool("no-confirm") {
				fmt.Printf("Security policy [%v] version [%v] CIDR [%v] "+
					"will be delete!\n",
					index, version, policy)
				fmt.Printf("Confirm [y/N]: ")
				var confirm string
				fmt.Scanf("%s", &confirm)
				if confirm != "y" && confirm != "Y" {
					return fmt.Errorf("abort")
				}
			}

			if cmdconfig.GetBool("dry-run") {
				logrus.Infof("dry-run set, skip")
				return nil
			}

			response, err := tcr.DeleteSecurityPolicy(index, policy, version)
			if err != nil {
				return fmt.Errorf(
					"DeleteMultipleSecurityPolicy failed: %w", err)
			}
			logrus.Debugf("%v", response.ToJsonString())
			logrus.Infof("Successfully remove %q from security policy", policy)
			return nil
		},
	}
	cc.cmd.Flags().StringP("config", "", utils.TAC_CONFIG_FILE, "config file")
	cc.cmd.Flags().StringP("ip", "", "", "IPv4 address (CIDR block) (required)")
	cc.cmd.Flags().Int64P("index", "", -1, "Policy index number (required)")
	cc.cmd.Flags().BoolP("no-confirm", "y", false, "Auto yes")
	cc.cmd.Flags().BoolP("dry-run", "", false, "Dry run, do not delete")

	return cc
}

func (cc *removeCmd) getCommand() *cobra.Command {
	return cc.cmd
}
