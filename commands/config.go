package commands

import (
	"encoding/json"
	"fmt"

	"github.com/cnrancher/tcr-access-control/pkg/cmdconfig"
	"github.com/cnrancher/tcr-access-control/pkg/config"
	"github.com/cnrancher/tcr-access-control/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type configCmd struct {
	baseCmd
}

func newConfigCmd() *configCmd {
	cc := &configCmd{}

	cc.baseCmd.cmd = &cobra.Command{
		Use:   "init",
		Short: "Setup Tencent Cloud Config",
		Example: `  tcr-access-control init \
	--registryID=<tcr-instance-id> \
	--language=zh-CN \
	--region=ap-guangzhou \
	--secretID=xxx \
	--secretKey=xxx`,
		RunE: func(cmd *cobra.Command, args []string) error {
			initializeFlagsConfig(cmd, cmdconfig.DefaultProvider)
			logrus.Debugf("Debug output enabled")
			if cmdconfig.GetString("config") == "" {
				return fmt.Errorf("config file not specified")
			}
			logrus.Debugf("config file: %v", cmdconfig.GetString("config"))
			cfg := config.Config{}
			if cmdconfig.GetString("secretID") != "" &&
				cmdconfig.GetString("secretKey") != "" &&
				cmdconfig.GetString("registryID") != "" {
				// Get config from command-line parameter
				cfg = config.Config{
					Language:  cmdconfig.GetString("language"),
					Region:    cmdconfig.GetString("region"),
					SecretID:  cmdconfig.GetString("secretID"),
					SecretKey: cmdconfig.GetString("secretKey"),
				}
				var err error
				cfg.SecretID, err = utils.EncryptAES(
					utils.AesEncryptKey, cfg.SecretID)
				if err != nil {
					return fmt.Errorf("Failed to encrypt secretID: %v", err)
				}
				cfg.SecretKey, err = utils.EncryptAES(
					utils.AesEncryptKey, cfg.SecretKey)
				if err != nil {
					return fmt.Errorf("Failed to encrypt secretKey: %v", err)
				}
				cfg.RegistryID = cmdconfig.GetString("registryID")
			} else {
				// Get config from user input
				logrus.Infof("Start init config:")
				fmt.Printf(
					"Default language (zh-CN/en-US) (default: en-US): ")
				fmt.Scanf("%s", &cfg.Language)
				switch cfg.Language {
				case "zh-CN":
				case "en-US":
				case "":
					logrus.Infof("Set language to default 'en-US'")
					cfg.Language = "en-US"
				default:
					logrus.Errorf(
						"Invalid language [%s], set to 'en-US'", cfg.Language)
					cfg.Language = "en-US"
				}

				fmt.Printf("Region (default: ap-guangzhou): ")
				fmt.Scanf("%s", &cfg.Region)
				if cfg.Region == "" {
					logrus.Infof("Set region to default 'ap-guangzhou'")
					cfg.Region = "ap-guangzhou"
				}

				fmt.Printf("SecretID: ")
				fmt.Scanf("%s", &cfg.SecretID)
				if cfg.SecretID == "" {
					return fmt.Errorf("secretID should not be empty")
				}
				var err error
				cfg.SecretID, err = utils.EncryptAES(
					utils.AesEncryptKey, cfg.SecretID)
				if err != nil {
					return fmt.Errorf("Failed to encrypt secretID: %v", err)
				}

				fmt.Printf("SecretKey: ")
				fmt.Scanf("%s", &cfg.SecretKey)
				if cfg.SecretKey == "" {
					return fmt.Errorf("secretKey should not be empty")
				}
				cfg.SecretKey, err = utils.EncryptAES(
					utils.AesEncryptKey, cfg.SecretKey)
				if err != nil {
					return fmt.Errorf("Failed to encrypt secretKey: %v", err)
				}

				fmt.Printf("TCR Instance ID: ")
				fmt.Scanf("%s", &cfg.RegistryID)
				if cfg.RegistryID == "" {
					return fmt.Errorf("TCR Instance ID should not be empty")
				}
			}

			b, _ := json.MarshalIndent(cfg, "", "  ")
			logrus.Debugf("config: %v", string(b))

			err := config.SaveConfig(&cfg, cmdconfig.GetString("config"))
			if err != nil {
				return err
			}
			logrus.Infof("Saved config to %q", cmdconfig.GetString("config"))
			return nil
		},
	}
	cc.cmd.Flags().StringP("language", "l", "en-US", "Laguage")
	cc.cmd.Flags().StringP("region", "r", "ap-guangzhou", "Region")
	cc.cmd.Flags().StringP("secretID", "", "", "secretID")
	cc.cmd.Flags().StringP("secretKey", "", "", "secretKey")
	cc.cmd.Flags().StringP("registryID", "", "", "registryID")
	cc.cmd.Flags().StringP("config", "", utils.TAC_CONFIG_FILE, "config file")

	return cc
}

func (cc *configCmd) getCommand() *cobra.Command {
	return cc.cmd
}
