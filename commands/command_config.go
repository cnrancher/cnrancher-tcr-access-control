package commands

import (
	"fmt"
	"strings"

	"github.com/cnrancher/tcr-access-control/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

func initializeFlagsConfig(cmd *cobra.Command, cfg config.Provider) {
	if cmd.Parent() != nil {
		initializeFlagsConfig(cmd.Parent(), cfg)
	}

	flags := cmd.Flags()
	cmd.Flags().VisitAll(func(f *flag.Flag) {
		setValueFromFlag(flags, f.Name, cfg, "")
	})

	if cfg.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func setValueFromFlag(flags *flag.FlagSet, key string, cfg config.Provider, targetKey string) {
	key = strings.TrimSpace(key)
	if flags.Lookup(key) != nil || flags.Changed(key) {
		f := flags.Lookup(key)
		configKey := key
		if targetKey != "" {
			configKey = targetKey
		}
		// Gotta love this API too.
		switch f.Value.Type() {
		case "bool":
			bv, _ := flags.GetBool(key)
			cfg.Set(configKey, bv)
		case "string":
			cfg.Set(configKey, f.Value.String())
		case "stringSlice":
			sv, _ := flags.GetStringSlice(key)
			cfg.Set(configKey, sv)
		case "int":
			iv, _ := flags.GetInt(key)
			cfg.Set(configKey, iv)
		case "int32":
			iv, _ := flags.GetInt32(key)
			cfg.Set(configKey, iv)
		case "int64":
			iv, _ := flags.GetInt64(key)
			cfg.Set(configKey, iv)
		default:
			panic(fmt.Sprintf("update switch with %s", f.Value.Type()))
		}
	}
}
