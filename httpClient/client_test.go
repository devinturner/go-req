package httpClient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		clientOut  *Client
	}{
		{"no header -> success", []string{}, &Client{Headers: map[string]string{}}},
    {
      "User-Agent=go-req/0.1.0 header -> success",
      []string{"User-Agent=go-req/0.1.0"},
      &Client{Headers: map[string]string{"User-Agent": "go-req/0.1.0"}},
    },
		{
			"User-Agent=go-req/0.1.0 + Content-Type=application/json -> success",
      []string{"User-Agent=go-req/0.1.0", "Content-Type=application/json"},
			&Client{
				Headers: map[string]string{
					"User-Agent": "go-req/0.1.0",
					"Content-Type": "application/json",
				},
			},
		},
	}

  for _, tt := range tests {
	  t.Run(tt.name, func(t *testing.T) {
		   assert.EqualValues(t, tt.clientOut, NewClient(tt.in))
	  })
  }
}
