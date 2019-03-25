package httpClient

import (
	"net/http"
)

// Get build an HTTP GET request and returns a Response and an error
func (c *Client) Get(uri string) (*Response, error) {
	cc := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range c.Headers {
		req.Header.Add(k, v)
	}

	raw, err := cc.Do(req)
	if err != nil {
		return nil, err
	}

	return decodeResponse(raw), nil
}

// decodeResponse takes in an http.Response object and returns a Response
func decodeResponse(raw *http.Response) *Response {
	resp := &Response{
		StatusCode: raw.StatusCode,
		Headers:    decodeHeaders(raw.Header),
		Body:       raw.Body,
	}
	defer raw.Body.Close()
	return resp
}

// decodeHeaders takes in an http.Header object and returns a map
func decodeHeaders(raw http.Header) map[string]string {
	headers := map[string]string{}
	for k, v := range raw {
		headers[k] = v[0]
	}
	return headers
}
