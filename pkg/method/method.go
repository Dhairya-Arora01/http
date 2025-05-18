// Package method contains utilities to parse and work with HTTP methods.
package method

// Method represents an HTTP method/verb.
type Method string

const (
	// GET represents the HTTP GET method.
	GET Method = "GET"

	// HEAD represents the HTTP HEAD method.
	HEAD Method = "HEAD"

	// POST represents the HTTP POST method.
	POST Method = "POST"

	// PUT represents the HTTP PUT method.
	PUT Method = "PUT"

	// DELETE represents the HTTP DELETE method.
	DELETE Method = "DELETE"

	// CONNECT represents the HTTP CONNECT method.
	CONNECT Method = "CONNECT"

	// OPTIONS represents the HTTP OPTIONS method.
	OPTIONS Method = "OPTIONS"

	// TRACE represents the HTTP TRACE method.
	TRACE Method = "TRACE"

	// PATCH represents the HTTP PATCH method.
	PATCH Method = "PATCH"
)
