package gists

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigFileName(t *testing.T) {
	filename := GetConfigFileName()
	assert.NotEmpty(t, filename)
}

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name         string
		filename     string
		wantErr      bool
		wantUsername string
		wantToken    string
	}{
		{
			name:         "Happy path",
			filename:     filepath.Join("testdata", "config.yaml"),
			wantErr:      false,
			wantUsername: "Curly",
			wantToken:    "woowoowoo",
		},
		{
			name:    "empty",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := LoadConfig(tt.filename)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, config)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, config)
				assert.Equal(t, tt.wantUsername, config.Username)
			}
		})
	}
}
