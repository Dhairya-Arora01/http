// Package header provides utilities for parsing, constructing, and manipulating
// HTTP headers in a structured and consistent manner.
package header

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// DateHeaderName represents the name of Date header.
	DateHeaderName = "Date"

	// ContentTypeHeaderName represents the name of content-type header.
	ContentTypeHeaderName = "content-type"

	// ContentLengthHeaderName represents the name of content-length header.
	ContentLengthHeaderName = "content-length"
)

// Header represents an HTTP header.
// An HTTP header is a key value pair that conveys request/response metadata.
type Header struct {
	Key   string // Key is a case-insensitive string.
	Value []string
}

// ContentLength checks if the specified header is a "Content-Length" header.
// If it is, it returns the specified value of content-length.
func (h *Header) ContentLength() (isContentLength bool, length int, err error) {
	if h.Key != ContentLengthHeaderName {
		return false, 0, nil
	}

	length, err = strconv.Atoi(h.Value[0])
	if err != nil {
		return true, 0, fmt.Errorf("connot parse the content-length into type int: %w", err)
	}

	return true, length, nil
}

// FromString constructs a Header from a line in the request/response metadata.
func FromString(headerLine string) (*Header, error) {
	keyVal := strings.SplitN(headerLine, ":", 2)
	if len(keyVal) != 2 {
		return nil, errors.New("malformed header-line")
	}

	// for multiple values in single header line, split the values on "," and remove whitespaces.
	valSlice := strings.Split(keyVal[1], ",")
	for i := range valSlice {
		valSlice[i] = strings.TrimSpace(valSlice[i])
	}

	return &Header{
		Key:   strings.ToLower(keyVal[0]), // the key is case insensitive, so make it lower case.
		Value: valSlice,
	}, nil
}

// New returns a new Header.
func New(key, value string) *Header {
	return &Header{
		Key:   key,
		Value: []string{value},
	}
}
