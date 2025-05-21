// Package router provide utilities to manage and handle routes.
package router

import (
	"bytes"
	"fmt"

	"github.com/Dhairya-Arora01/http/pkg/method"
	"github.com/Dhairya-Arora01/http/pkg/request"
	"github.com/Dhairya-Arora01/http/pkg/response"
	"github.com/Dhairya-Arora01/http/pkg/status"
)

// HandlerFunc defines the signature for a request handler.
// It processes an incoming request for a given URL and set of HTTP methods,
// and returns the appropriate response.
type HandlerFunc func(req *request.Request) (*response.Response, error)

// Router stores a collection of registered handler functions.
// It is responsible for routing incoming requests to the appropriate handler based on the URL and method.
type Router struct {
	routes map[string]HandlerFunc
}

// New returns a new instance of Router.
func New() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}

// routeKey returns a string compose of URL and method.
// This key is used to map the handler function to a specific route in the Router.
func routeKey(url string, method method.Method) string {
	return fmt.Sprintf("%s %s", method, url)
}

// RegisterHandler registers a handler function with the router.
func (r *Router) RegisterHandler(url string, methods []method.Method, f HandlerFunc) {
	for _, method := range methods {
		rKey := routeKey(url, method)
		r.routes[rKey] = f
	}
}

// Handle routes the incoming request to the appropriate HandlerFunc based on the request's URL and method.
// It invokes the matched handler and returns the corresponding response.
// If no matching handler is found, it should return a response with status 404 Not Found.
func (r *Router) Handle(req *request.Request) *response.Response {
	rKey := routeKey(req.URL.Path, req.Method)
	f, ok := r.routes[rKey]
	// if handler function for the route is not found return a 404 Not Found error.
	if !ok {
		return &response.Response{
			Status: status.NotFound,
		}
	}

	res, err := f(req)
	if err != nil {
		return &response.Response{
			Status: status.InternalServerError,
			Body:   bytes.NewBufferString(err.Error()),
		}
	}

	return res
}
