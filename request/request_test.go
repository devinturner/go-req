package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildRequest(t *testing.T) {
	tests := []struct {
		name        string
		in          Request
		uri         string
		out         http.Request
		errExpected bool
	}{
		{"nil returns successfully", nil, "", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := buildRequest(tt.in, tt.uri)
			if tt.errExpected {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.out, actual)
		})
	}
}
