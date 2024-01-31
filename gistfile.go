package gists

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GistFile struct {
	Filename string `json:"filename,omitempty"`
	Type     string `json:"type,omitempty"`
	Language string `json:"language,omitempty"`
	RawURL   string `json:"raw_url,omitempty"`
	Size     int    `json:"size,omitempty"`
}

func NewGistFile(jsonstr string) (*GistFile, error) {
	gf := new(GistFile)
	jsonbytes := []byte(jsonstr)
	err := json.Unmarshal(jsonbytes, gf)
	return gf, err
}

func (gf GistFile) String() string {
	parts := make([]string, 0)
	parts = append(parts, fmt.Sprintf("Filename: %q", gf.Filename))
	return strings.Join(parts, "\n")
}

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