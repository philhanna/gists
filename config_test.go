package gists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfigFileName(t *testing.T) {
	filename := GetConfigFileName()
	assert.NotEmpty(t, filename)
}
