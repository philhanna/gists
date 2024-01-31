package gists

import (
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreateDatabase(t *testing.T) {
	tests := []struct {
		name          string
		inputFileName string
		dbFileName    string
	}{
		{
			name:          "Happy path",
			inputFileName: filepath.Join("testdata", "gistpage1.json"),
			dbFileName:    "/tmp/gists.db",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := os.ReadFile(tt.inputFileName)
			assert.Nil(t, err)
			jsonstr := string(body)
			ch := GetGists(jsonstr)
			err = CreateDatabase(tt.dbFileName, ch)
			assert.Nil(t, err)
		})
	}
}
