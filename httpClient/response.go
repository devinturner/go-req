package httpClient

import (
	"encoding/json"
	"fmt"
	"io"
)

// Response defines the expected json for marshalling
type Response struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       io.ReadCloser     `json:"body"`
}

// ToJSON marshals and prints the Response
func (resp *Response) ToJSON() error {
	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
