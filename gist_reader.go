package gists

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	URL_PREFIX = "https://api.github.com/users"
	URL_SUFFIX = "gists?per_page=10"
)

func GetGists() {
	// Load the configuration
	config, err := LoadConfig(GetConfigFileName())
	if err != nil {
		log.Fatal(err)
	}

	// Create the request
	url := fmt.Sprintf("%s/%s/%s", URL_PREFIX, config.Username, URL_SUFFIX)
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Token))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	// Make the http request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Get the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print it
	fmt.Println(string(body))

	// Get the links
	for key, values := range resp.Header {
		fmt.Println("Header:", key)
		for _, value := range values {
			fmt.Println("Value:", value)
		}
	}
}
