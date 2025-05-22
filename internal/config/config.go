package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config holds the application's configuration
type Config struct {
	ProjectRoots []string `json:"project_roots"`
}

const (
	configDirName  = "ep"
	configFileName = "config.json"
)

// getConfigPath returns the full path to the configuration file.
func getConfigPath() (string, string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", "", fmt.Errorf("error getting user config directory: %w", err)
	}
	appConfigDirPath := filepath.Join(userConfigDir, configDirName)
	configFilePath := filepath.Join(appConfigDirPath, configFileName)
	return appConfigDirPath, configFilePath, nil
}

// LoadConfig reads the configuration from the file or returns a default config if not found.
func LoadConfig() (*Config, error) {
	_, configFilePath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	var cfg Config
	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Config file not found, return a default empty config
			// The 'add' command will create it when a path is first added.
			fmt.Printf("Config file not found at %s. A new one will be created when you add a path.\n", configFilePath)
			return &Config{ProjectRoots: []string{}}, nil
		}
		return nil, fmt.Errorf("error reading config file %s: %w", configFilePath, err)
	}

	if err := json.Unmarshal(configData, &cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config data from %s: %w", configFilePath, err)
	}
	return &cfg, nil
}

// SaveConfig writes the current configuration to the file.
func SaveConfig(cfg *Config) error {
	appConfigDirPath, configFilePath, err := getConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(appConfigDirPath, 0750); err != nil {
		return fmt.Errorf("error creating config directory %s: %w", appConfigDirPath, err)
	}

	updatedConfigData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling config to JSON: %w", err)
	}

	if err := os.WriteFile(configFilePath, updatedConfigData, 0640); err != nil {
		return fmt.Errorf("error writing config file %s: %w", configFilePath, err)
	}
	fmt.Printf("Configuration updated successfully at %s\n", configFilePath)
	return nil
}
