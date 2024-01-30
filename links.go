package gists

import (
	"regexp"
	"strings"
)

type Link struct {
	URL string
	Rel string
}

func SplitLinks(s string) []Link {
	links := make([]Link, 0)
	re := regexp.MustCompile(`<(.*)>; rel="(.*)"`)
	for _, x := range strings.Split(s, ",") {
		m := re.FindAllStringSubmatch(x, -1)
		for _, matches := range m {
			links = append(links, Link{
				URL: matches[1],
				Rel: matches[2],
			})
		}
	}
	return links
}
