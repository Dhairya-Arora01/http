package response

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/Dhairya-Arora01/http/pkg/header"
	"github.com/Dhairya-Arora01/http/pkg/status"
)

const (
	// HTTPProtocol represents the HTTP protocol 1.1 i.e HTTP/1.1.
	HTTPProtocol = "HTTP/1.1"

	// TimeFormat represents the time format accepted by the HTTP protocol.
	TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

	// DateHeaderName represents the name of Date header.
	DateHeaderName = "Date"
)

// Request represents an HTTP request.
type Response struct {
	Status  status.Status
	Headers []*header.Header
	Body    io.Reader
}

// New returns a new Response object.
func New(status status.Status, headers []*header.Header, body io.Reader) *Response {
	responseHeaders := []*header.Header{}

	dateHeader := &header.Header{
		Key: DateHeaderName,
		Value: []string{
			time.Now().UTC().Format(TimeFormat),
		},
	}

	responseHeaders = append(responseHeaders, dateHeader)
	responseHeaders = append(responseHeaders, headers...)

	return &Response{
		Status:  status,
		Headers: responseHeaders,
		Body:    body,
	}
}

// Write writes the response to the request.
func (r *Response) Write(conn net.Conn) error {
	startLine := fmt.Sprintf("%s %d %s\r\n", HTTPProtocol, r.Status.Code, r.Status.Text)

	var headers string
	for _, header := range r.Headers {
		var headerValue string
		for i := range header.Value {
			if i == 0 {
				headerValue = header.Value[i]
			} else {
				headerValue += fmt.Sprintf(", %s", header.Value[i])
			}
		}

		headers += fmt.Sprintf("%s: %s\r\n", header.Key, headerValue)
	}

	var (
		body []byte
		err  error
	)

	if r.Body != nil {
		body, err = io.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}
	}

	res := startLine + headers + "\r\n" + string(body)
	if _, err = conn.Write([]byte(res)); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}
