package request

import (
	"net/http"
	"strings"
	"time"
)

type Request struct {
	Headers map[string]string
}

func NewRequest(args []string) *Request {
	headers := map[string]string{}

	for _, header := range args {
		h := strings.Split(header, "=")
		headers[h[0]] = h[1]
	}

	return &Request{Headers: headers}
}

func (r *Request) Get(uri string) (*http.Response, error) {
	req, err := buildRequest(r, uri)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	raw, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func buildRequest(r *Request, uri string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	return req, nil
}
