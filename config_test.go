package gists

import (
	"fmt"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestGetConfigFileName(t *testing.T) {
	filename := GetConfigFileName()
	fmt.Println(filename)
}
