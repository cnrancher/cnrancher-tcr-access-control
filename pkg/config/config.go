package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Language   string `json:"language" yaml:"language"`
	Region     string `json:"region" yaml:"region"`
	SecretID   string `json:"secretID" yaml:"secretID"`
	SecretKey  string `json:"secretKey" yaml:"secretKey"`
	RegistryID string `json:"registryID" yaml:"registryID"`
}

// LoadConfig loads the config from file path
func LoadConfig(filename string) (*Config, error) {
	logrus.Debugf("Load config from %q", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("LoadConfig: %w", err)
	}
	return config, nil
}

// SaveConfig saves config into file
func SaveConfig(config *Config, filename string) error {
	err := os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		return fmt.Errorf("SaveConfig: %w", err)
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("SaveConfig: %w", err)
	}
	return os.WriteFile(filename, data, 0644)
}
