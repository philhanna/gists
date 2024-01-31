package gists

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreateDatabase(t *testing.T) {
	tests := []struct {
		name          string
		dbFileName    string
	}{
		{
			name:          "Happy path",
			dbFileName:    "/tmp/gists.db",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateDatabase(tt.dbFileName)
			assert.Nil(t, err)
		})
	}
}
