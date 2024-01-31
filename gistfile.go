package gists

import (
	"fmt"
	"strings"
)

type GistFile struct {
	Filename string `json:"filename,omitempty"`
	Type     string `json:"type,omitempty"`
	Language string `json:"language,omitempty"`
	RawURL   string `json:"raw_url,omitempty"`
	Size     int    `json:"size,omitempty"`
}

func (gf GistFile) String() string {
	parts := make([]string, 0)
	parts = append(parts, fmt.Sprintf("Filename: %q", gf.Filename))
	return strings.Join(parts, "\n")
}