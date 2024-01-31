package gists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGistFile_GetContents(t *testing.T) {
	tests := []struct {
		name string
		json string
		want string
	}{
		{
			name: "Happy path",
			json: `{
				"filename": "generator.go",
				"type": "text/plain",
				"language": "Go",
				"raw_url": "https://gist.githubusercontent.com/philhanna/220b3d5a2063c8942dd124b87f55e154/raw/cfab88b1c9868d80a48ac413cf5a48fc1005f49d/generator.go",
				"size": 460
			}`,
			want: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gf, err := NewGistFile(tt.json)
			assert.Nil(t, err)
			have, err := gf.GetContents()
			assert.Nil(t, err)
			fmt.Println(string(have))
		})
	}
}
