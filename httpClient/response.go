package httpClient

import (
	"encoding/json"
	"fmt"
	"io"
)

type Response struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       io.ReadCloser     `json:"body"`
}

func (resp *Response) ToJSON() error {
	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
