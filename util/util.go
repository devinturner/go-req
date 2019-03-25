package util

import (
	"errors"
	"net/url"
)

const argCount = 1

// GetURI validates and returns the URI from the passed in arguments
func GetURI(args []string) (string, error) {
	if len(args) != argCount {
		return "", errors.New("invalid number or args")
	}

	uri, err := url.ParseRequestURI(args[0])
	if err != nil {
		return "", errors.New("invalid url")
	}
	return uri.String(), nil
}
