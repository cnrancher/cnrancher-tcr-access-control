package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cnrancher/tcr-access-control/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

var (
	SECRET_ID_ENV_NAME   = "TENCENT_CLOUD_SECRET_ID"
	SECRET_KEY_ENV_NAME  = "TENCENT_CLOUD_SECRET_KEY"
	REGION_ENV_NAME      = "TENCENT_CLOUD_REGION"
	REGISTRY_ID_ENV_NAME = "TENCENT_CLOUD_REGISTRY_ID"

	// TAC_CONFIG_FILE_HOME is $HOME/.tcr_access_control.yaml
	TAC_CONFIG_FILE_HOME = filepath.FromSlash(os.Getenv("HOME") + "/.tcr_access_control.yaml")

	// TAC_CONFIG_FILE_GLOBAL is /etc/tcr-access-control/config.yaml
	TAC_CONFIG_FILE_GLOBAL = filepath.FromSlash("/etc/tcr-access-control/config.yaml")
)

var (
	ClientProfile *profile.ClientProfile
	Credential    *common.Credential
	Config        *config.Config

	initialized bool
)

// Init initializes the credential key and client profile
func Init(configPath string) error {
	if initialized && ClientProfile != nil &&
		Credential != nil && Config != nil {
		return nil
	}
	logrus.Debugf("Start init utils config")

	configInitErrMsg := fmt.Sprintf("execute '%s init' first", os.Args[0])

	var err error
	if configPath != "" {
		Config, err = config.LoadConfig(configPath)
		if err != nil {
			return fmt.Errorf("LoadConfig: %w", err)
		}
	} else {
		// try read config from $HOME
		Config, err = config.LoadConfig(TAC_CONFIG_FILE_HOME)
		if err != nil {
			if !os.IsNotExist(err) {
				return fmt.Errorf("LoadConfig: %w", err)
			}
			// try read config from /etc/tcr-access-control/
			Config, err = config.LoadConfig(TAC_CONFIG_FILE_GLOBAL)
			if err != nil {
				return fmt.Errorf("LoadConfig: %w", err)
			}
		}
	}

	if Config.SecretKey == "" || Config.SecretID == "" {
		return fmt.Errorf("credential not set in config, " + configInitErrMsg)
	}
	if Config.Region == "" {
		return fmt.Errorf("region not set in config, " + configInitErrMsg)
	}
	if Config.RegistryID == "" {
		return fmt.Errorf("registryID not set in config, " + configInitErrMsg)
	}

	Config.SecretID, err = DecryptAES(AesEncryptKey, Config.SecretID)
	if err != nil {
		return fmt.Errorf("failed to decrypt secretID: %w", err)
	}
	Config.SecretKey, err = DecryptAES(AesEncryptKey, Config.SecretKey)
	if err != nil {
		return fmt.Errorf("failed to decrypt secretKey: %w", err)
	}

	logrus.Debugf("Set Language: %v", Config.Language)
	logrus.Debugf("Set Region: %v", Config.Region)
	logrus.Debugf("Set TCR Instance ID: %v", Config.RegistryID)
	if len(Config.SecretID) < 8 {
		return fmt.Errorf("invalid secretID length, " + configInitErrMsg)
	}
	logrus.Debugf("Set SecretID: [%s*****]", Config.SecretID[:8])

	ClientProfile = profile.NewClientProfile()
	ClientProfile.Language = Config.Language
	Credential = common.NewCredential(
		Config.SecretID,
		Config.SecretKey,
	)

	initialized = true
	logrus.Debugf("Finished init utils config")
	return nil
}

type valueTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 |
		~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string | ~bool |
		[]string
}

// Pointer gets the pointer of the variable.
func Pointer[T valueTypes](i T) *T {
	return &i
}

// A safe function to get the value from the pointer.
func Value[T valueTypes](p *T) T {
	if p == nil {
		return *new(T)
	}
	return *p
}
