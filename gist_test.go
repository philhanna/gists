package gists

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGists(t *testing.T) {
	filename := filepath.Join("testdata", "gistpage1.json")
	body, err := os.ReadFile(filename)
	assert.Nil(t, err)
	jsonstr := string(body)
	for gist := range GetGists(jsonstr) {
		fmt.Printf("%v\n", gist)
	}
}
