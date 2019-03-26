package response

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeHeaders(t *testing.T) {
	tests := []struct {
		name string
		in   http.Header
		out  map[string]string
	}{
		{"nil returns nil", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, decodeHeaders(tt.in))
		})
	}
}

func TestDecodeBody(t *testing.T) {
	tests := []struct {
		name string
		in   io.ReadCloser
		out  []byte
	}{
		{"nil returns nil", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, decodeBody(tt.in))
		})
	}
}
