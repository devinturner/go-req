package httpClient

import (
	"strings"
)

type Client struct {
	Headers map[string]string
}

func NewClient(args []string) *Client {
	headers := map[string]string{}

	for _, header := range args {
		h := strings.Split(header, "=")
		headers[h[0]] = h[1]
	}

	return &Client{Headers: headers}
}
