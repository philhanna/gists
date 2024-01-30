package gists

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := LoadConfig(tt.filename)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.wantUsername, config.Username)
			assert.Equal(t, tt.wantToken, config.Token)
		})
	}
}
