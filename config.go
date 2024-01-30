package gists

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Configuration contains the Github username and token
type Configuration struct {
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
}

const (
	PACKAGE_NAME     = "get-gists"
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
	config := new(Configuration)
	body, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(body, config)
	return config, err
}
