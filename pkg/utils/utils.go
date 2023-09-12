package utils

import (
	"fmt"
	"os"
	"path"

	tcr_config "github.com/cnrancher/tcr-access-control/pkg/tcr-config"
	"github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

var (
	SECRET_ID_ENV_NAME  = "TENCENT_CLOUD_SECRET_ID"
	SECRET_KEY_ENV_NAME = "TENCENT_CLOUD_SECRET_KEY"
	REGION_ENV_NAME     = "TENCENT_CLOUD_REGION"
	TAC_CONFIG_FILE     = path.Join(
		os.Getenv("HOME"),
		".tcr_access_controll.yaml",
	)
)

var (
	ClientProfile *profile.ClientProfile
	Credential    *common.Credential

	SecretID  = os.Getenv(SECRET_ID_ENV_NAME)
	SecretKey = os.Getenv(SECRET_KEY_ENV_NAME)
	Region    = os.Getenv(REGION_ENV_NAME)

	Config *tcr_config.Config

	initialized bool
)

// Init initializes the credential key and client profile
func Init(configPath string) error {
	if initialized && ClientProfile != nil &&
		Credential != nil && Config != nil {
		return nil
	}
	logrus.Debugf("Start init config")

	// load config
	var err error
	Config, err = tcr_config.LoadConfig(configPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("utils.Init: %w", err)
		}
		Config = &tcr_config.Config{}
	}
	if Region == "" {
		Region = Config.Region
	}
	if SecretID == "" || SecretKey == "" {
		if Config.SecretID != "" {
			SecretID, err = DecryptAES(AesEncryptKey, Config.SecretID)
			if err != nil {
				return fmt.Errorf("failed to decrypt secretID: %w", err)
			}
		}
		if Config.SecretKey != "" {
			SecretKey, err = DecryptAES(AesEncryptKey, Config.SecretKey)
			if err != nil {
				return fmt.Errorf("failed to decrypt secretKey: %w", err)
			}
		}
	}

	if SecretID == "" || SecretKey == "" {
		return fmt.Errorf("%s & %s env not set",
			SECRET_ID_ENV_NAME, SECRET_KEY_ENV_NAME)
	}
	if Region == "" {
		return fmt.Errorf("%s env not set", REGION_ENV_NAME)
	}

	logrus.Debugf("Set Language: %v", Config.Language)
	logrus.Debugf("Set Region: %v", Config.Region)
	logrus.Debugf("Set TCR Instance ID: %v", Config.InstanceID)

	ClientProfile = profile.NewClientProfile()
	ClientProfile.Language = Config.Language
	Credential = common.NewCredential(
		SecretID,
		SecretKey,
	)

	initialized = true
	logrus.Debugf("Finished init config")
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
