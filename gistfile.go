package gists

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GistFile represents a single file in a gist, not including the content,
// which is pointed to by RawURL.
type GistFile struct {
	Filename string `json:"filename,omitempty"`
	Type     string `json:"type,omitempty"`
	Language string `json:"language,omitempty"`
	RawURL   string `json:"raw_url,omitempty"`
	Size     int    `json:"size,omitempty"`
}

// GetContents retrieves the actual file contents from GitHub.
func (gf GistFile) GetContents() ([]byte, error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, gf.RawURL, nil)
	if err != nil {
		return nil, err
	}
	AddHeaders(req)

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Get the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// NewGistFile creates a new GistFile structure from its JSON
// representation.
func NewGistFile(jsonstr string) (*GistFile, error) {
	gf := new(GistFile)
	jsonbytes := []byte(jsonstr)
	err := json.Unmarshal(jsonbytes, gf)
	return gf, err
}

// String returns a string representation of a GistFile.
func (gf GistFile) String() string {
	parts := make([]string, 0)
	parts = append(parts, fmt.Sprintf("Filename: %q", gf.Filename))
	return strings.Join(parts, "\n")
}
