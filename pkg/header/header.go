// Package header provides utilities for parsing, constructing, and manipulating
// HTTP headers in a structured and consistent manner.
package header

import (
	"errors"
	"strings"
)

// Header represents an HTTP header.
// An HTTP header is a key value pair that conveys request/response metadata.
type Header struct {
	Key   string // Key is a case-insensitive string.
	Value string
}

// FromString constructs a Header from a line in the request/response metadata.
func FromString(headerLine string) (*Header, error) {
	keyVal := strings.SplitN(headerLine, ":", 2)
	if len(keyVal) != 2 {
		return nil, errors.New("malformed header-line")
	}

	return &Header{
		Key:   strings.ToLower(keyVal[0]), // the key is case insensitive, so make it lower case.
		Value: strings.TrimSpace(keyVal[1]),
	}, nil
}
