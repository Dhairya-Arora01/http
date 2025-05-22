package response

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"

	"github.com/Dhairya-Arora01/http/pkg/header"
	"github.com/Dhairya-Arora01/http/pkg/status"
)

const (
	// HTTPProtocol represents the HTTP protocol 1.1 i.e HTTP/1.1.
	HTTPProtocol = "HTTP/1.1"

	// TimeFormat represents the time format accepted by the HTTP protocol.
	TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
)

// Request represents an HTTP request.
type Response struct {
	Status  status.Status
	Headers []*header.Header
	Body    io.Reader
}

// New returns a new Response object.
func New(status status.Status, contentType ContentType, headers []*header.Header, body io.Reader) *Response {
	responseHeaders := []*header.Header{}

	dateHeader := header.New(header.DateHeaderName, time.Now().UTC().Format(TimeFormat))
	contentTypeHeader := header.New(header.ContentTypeHeaderName, contentType.String())

	responseHeaders = append(responseHeaders, dateHeader, contentTypeHeader)
	responseHeaders = append(responseHeaders, headers...)

	return &Response{
		Status:  status,
		Headers: responseHeaders,
		Body:    body,
	}
}

// Write writes the response to the request.
func (r *Response) Write(conn net.Conn) error {
	var (
		buf           bytes.Buffer
		contentLength int64
		err           error
	)

	if r.Body != nil {
		contentLength, err = buf.ReadFrom(r.Body)
		if err != nil {
			contentLength = 0
		}
	}

	contentLengthHeader := header.New(header.ContentLengthHeaderName, strconv.FormatInt(contentLength, 10))
	r.Headers = append(r.Headers, contentLengthHeader)

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

	res := startLine + headers + "\r\n" + string(buf.Bytes())
	if _, err = conn.Write([]byte(res)); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}
