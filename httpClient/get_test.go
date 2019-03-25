package httpClient

import (
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
		{"nil returns nil", nil, map[string]string{}},
		{"empty returns empty", http.Header{}, map[string]string{}},
		{"non-empty returns valid", http.Header{"test": []string{"abc"}}, map[string]string{"test": "abc"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, tt.out, decodeHeaders(tt.in))
		})
	}
}
