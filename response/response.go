package response

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       []byte            `json:"body"`
}

func ParseResponse(resp *http.Response) *Response {
	return &Response{
		StatusCode: resp.StatusCode,
		Body:       decodeBody(resp.Body),
		Headers:    decodeHeaders(resp.Header),
	}
}

// decodeHeaders takes in an http.Header object and returns a map
func decodeHeaders(raw http.Header) map[string]string {
	headers := map[string]string{}
	for k, v := range raw {
		headers[k] = v[0]
	}
	return headers
}

func decodeBody(raw io.ReadCloser) []byte {
	body, err := ioutil.ReadAll(raw)
	if err != nil {
		panic(err)
	}
	raw.Close()
	return body
}
