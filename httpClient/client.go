package httpClient

import (
	"strings"
)

// Client struct defines an HTTP client with a map of headers
type Client struct {
	Headers map[string]string
}

// NewClient returns a Client with the headers passed in as an array of strings
// where key and value are separated by '='
func NewClient(args []string) *Client {
	headers := map[string]string{}

	for _, header := range args {
		h := strings.Split(header, "=")
		headers[h[0]] = h[1]
	}

	return &Client{Headers: headers}
}
