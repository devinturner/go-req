package util

import (
	"errors"
	"net/url"
)

const ARG_COUNT = 1

func GetURI(args []string) (string, error) {
	if len(args) != ARG_COUNT {
		return "", errors.New("invalid number or args")
	}

	uri, err := url.ParseRequestURI(args[0])
	if err != nil {
		return "", errors.New("invalid url")
	}
	return uri.String(), nil
}
