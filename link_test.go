package gists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitLinks(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []Link
	}{
		{
			name: "Happy path",
			s:    `<https://api.github.com/user/3708685/gists?per_page=10&page=2>; rel="next", <https://api.github.com/user/3708685/gists?per_page=10&page=6>; rel="last"`,
			want: []Link{
				{
					URL: `https://api.github.com/user/3708685/gists?per_page=10&page=2`,
					Rel: `next`,
				},
				{
					URL: `https://api.github.com/user/3708685/gists?per_page=10&page=6`,
					Rel: `last`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := SplitLinks(tt.s)
			assert.Equal(t, tt.want, have)
		})
	}
}
