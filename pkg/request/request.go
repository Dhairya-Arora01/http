// Package request offers tools for building, parsing, and analyzing HTTP requests.
package request

import (
	"io"

	"github.com/Dhairya-Arora01/http/pkg/header"
	"github.com/Dhairya-Arora01/http/pkg/method"
)

type Request struct {
	Method  method.Method
	URL     string
	Headers []*header.Header
	Body    io.Reader
}
