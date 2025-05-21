// Package request offers tools for building, parsing, and analyzing HTTP requests.
package request

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/url"
	"strings"

	"github.com/Dhairya-Arora01/http/pkg/header"
	"github.com/Dhairya-Arora01/http/pkg/method"
)

// Request represents an HTTP request.
type Request struct {
	Method  method.Method
	URL     *url.URL
	Headers []*header.Header
	Body    io.Reader
}

// New constructs a new Request object by reading the HTTP connection.
func New(conn net.Conn) (*Request, error) {
	reader := bufio.NewReader(conn)

	// Read start-line for the request.
	// <method> <request-target> <protocol>.
	startLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("no start-line of the form <method> <request-target> <protocol>: %w", err)
	}

	verb, rawURL, _, err := extractRequestLineParam(startLine)
	if err != nil {
		return nil, fmt.Errorf("invalid start-line: %w", err)
	}

	reqURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	var (
		headers       []*header.Header
		contentLength int
	)

	// Read Headers.
	for {
		headerLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read header: %w", err)
		}

		// headers span till the empty line that seggregates the headers with the request body.
		if headerLine == "\r\n" {
			break
		}

		reqHeader, err := header.FromString(headerLine)
		if err != nil {
			return nil, fmt.Errorf("failed to read header: %w", err)
		}

		isContentLength, length, err := reqHeader.ContentLength()
		if err != nil {
			return nil, fmt.Errorf("failed to extract the content-length header: %w", err)
		}

		if isContentLength {
			contentLength = length
		}

		headers = append(headers, reqHeader)
	}

	return &Request{
		Method:  verb,
		URL:     reqURL,
		Headers: headers,
		Body:    io.LimitReader(reader, int64(contentLength)),
	}, nil
}

func extractRequestLineParam(requestLine string) (verb method.Method, target string, protocol string, err error) {
	paramsSlice := strings.Split(requestLine, " ")
	if len(paramsSlice) < 3 {
		return method.Method(""), "", "", fmt.Errorf("request line with %d params, need to have 3 params of the form <method> <request-target> <protocol>", len(paramsSlice))
	}

	return method.Method(paramsSlice[0]), paramsSlice[1], paramsSlice[2], nil
}
