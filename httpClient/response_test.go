package httpClient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJSON(t *testing.T) {
	tests := []struct {
		name        string
		resp        *Response
		errExpected bool
	}{
		{"nil -> success", nil, false},
    {"empty -> success", &Response{}, false},
    {"partial -> success", &Response{StatusCode:200}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.resp.ToJSON()
			if tt.errExpected {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
