package tcr_config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Language   string `json:"language" yaml:"language"`
	Region     string `json:"region" yaml:"region"`
	SecretID   string `json:"secretID" yaml:"secretID"`
	SecretKey  string `json:"secretKey" yaml:"secretKey"`
	InstanceID string `json:"instanceID" yaml:"instanceID"`
}

// LoadConfig loads the config from file path
func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("LoadConfig: %w", err)
	}
	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("LoadConfig: %w", err)
	}
	return config, nil
}

// SaveConfig saves config into file
// (this method only use in test purpose)
func SaveConfig(config *Config, filename string) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("SaveConfig: %w", err)
	}
	return os.WriteFile(filename, data, 0644)
}
