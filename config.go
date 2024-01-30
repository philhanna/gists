package gists

import (
	"os"
	"path/filepath"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Configuration struct {
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
}

func GetConfigFileName() string {
	configDir, _ := os.UserConfigDir()
	filename := filepath.Join(configDir, "get-config", "config.yaml")
	return filename
}
