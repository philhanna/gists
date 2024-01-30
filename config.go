package gists

import (
	"errors"
	"os"
	"path/filepath"
)

// Configuration contains the Github username and token
type Configuration struct {
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
}

const (
	PACKAGE_NAME     = "get-config"
	CONFIG_FILE_NAME = "config.yaml"
)

// GetConfigFileName returns the name of the configuration file for this
// user.
func GetConfigFileName() string {
	configDir, _ := os.UserConfigDir()
	filename := filepath.Join(configDir, PACKAGE_NAME, CONFIG_FILE_NAME)
	return filename
}

// LoadConfig reads the specified file and produces a configuration
// structure from it by parsing it as YAML.
func LoadConfig(filename string) (*Configuration, error) {
	return nil, errors.New("dummy error message")
}
