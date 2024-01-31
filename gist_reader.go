package gists

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	URL_PREFIX       = "https://api.github.com/users"
	URL_SUFFIX       = "gists?per_page=10"
	LINK_HEADER_NAME = "Link"
	REL_NEXT_NAME    = "next"
)

// AddHeaders adds the required authorization headers to the HTTP
// request.
func AddHeaders(req *http.Request) {
	bearerToken := fmt.Sprintf("Bearer %s", Config.Token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", bearerToken)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
}

// GetGistPages downloads pages of gists from GitHub, sending them back as
// strings of JSON arrays.
func GetGistPages() chan string {
	ch := make(chan string, 10)
	go func() {
		defer close(ch)

		// Create the URL for the initial request
		url := MakeInitialURL()

		// Get all pages, starting with this URL
		for url != "" {
			client := new(http.Client)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				log.Fatal(err)
			}
			AddHeaders(req)

			// Make the request
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			// Get the body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			// Send it through the output channel
			ch <- string(body)

			// Get the URL for the next page (or blank, if there are no
			// more)
			url = MakeNextURL(resp)
		}
	}()
	return ch

}

// MakeInitialURL creates the URL for the first page of the gists
func MakeInitialURL() string {
	url := fmt.Sprintf("%s/%s/%s", URL_PREFIX, Config.Username, URL_SUFFIX)
	return url
}

// MakeNextURL finds the URL for the next page of the gists, or blank,
// if there are no more pages.
func MakeNextURL(resp *http.Response) string {
	url := ""
	for name, values := range resp.Header {
		if name == LINK_HEADER_NAME {
			for _, value := range values {
				links := SplitLinks(value)
				for _, link := range links {
					if link.Rel == REL_NEXT_NAME {
						url = link.URL
					}
				}
			}
		}
	}
	return url
}