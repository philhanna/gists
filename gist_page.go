package gists

import (
	"encoding/json"
	"log"
)

type Gist struct {
	URL string `json:"url,omitempty"`
}

// GetGists accepts a string containing a JSON array of gists
// and passes the individual gists back through a channel.
func GetGists(array string) chan Gist {
	ch := make(chan Gist, 10)
	go func() {
		defer close(ch)
		var gists []Gist
		err := json.Unmarshal([]byte(array), &gists)
		if err != nil {
			log.Fatal(err)
		}
		for _, gist := range gists {
			ch <- gist
		}
	}()
	return ch
}
