package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetURI(t *testing.T) {
	tests := []struct {
		name        string
		in          []string
		out         string
		errExpected bool
	}{
		{"no args -> fail", []string{}, "", true},
		{"single invalid -> fail", []string{"invalid"}, "", true},
		{"multiple invalid -> fail", []string{"invalid", "invalid"}, "", true},
		{"multiple valid -> fail", []string{"invalid", "invalid"}, "", true},
		{"single valid http -> success", []string{"http://example.org"}, "http://example.org", false},
		{"single valid custom -> success", []string{"custom://example.org"}, "custom://example.org", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetURI(tt.in)
			assert.Equal(t, tt.out, actual)
			if !tt.errExpected {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
