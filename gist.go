package gists

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Gist struct {
	ID          string              `json:"id,omitempty"`
	URL         string              `json:"url,omitempty"`
	Description string              `json:"description,omitempty"`
	CreatedAt   string              `json:"created_at"`
	Files       map[string]GistFile `json:"files,omitempty"`
}

func (g Gist) String() string {
	parts := make([]string, 0)
	parts = append(parts, fmt.Sprintf("ID: %q", g.ID))
	parts = append(parts, fmt.Sprintf("Description: %q", g.Description))
	for _, gistFile := range g.Files {
		parts = append(parts, fmt.Sprintf("File:\n\t%s", gistFile))
	}
	return strings.Join(parts, "\n")

}

// GetGists accepts a string containing a JSON array of gists
// and passes the individual gists back through a channel.
func GetGists(jsonstr string) chan Gist {
	ch := make(chan Gist, 10)
	go func() {
		defer close(ch)
		var gists []Gist
		err := json.Unmarshal([]byte(jsonstr), &gists)
		if err != nil {
			log.Fatal(err)
		}
		for _, gist := range gists {
			ch <- gist
		}
	}()
	return ch
}
