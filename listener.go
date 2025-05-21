// Package http provides the HTTP server implementation.
package http

import (
	"fmt"
	"net"

	"github.com/Dhairya-Arora01/http/pkg/port"
	"github.com/Dhairya-Arora01/http/pkg/request"
	"github.com/Dhairya-Arora01/http/pkg/router"
	"github.com/go-logr/logr"
)

// Start listens and accept connections at the specified port.
func Start(log logr.Logger, p int, pathRouter *router.Router) error {
	listener, err := net.Listen("tcp", port.ToString(p))
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listener.Addr().String(), err)
	}

	defer listener.Close()

	log.Info("Listening On", "address", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error(err, "Failed to accept connection")
			continue
		}

		if err := handleConnection(conn, pathRouter); err != nil {
			log.Error(err, "Failed to handle connection")
		}
	}
}

func handleConnection(conn net.Conn, pathRouter *router.Router) error {
	defer conn.Close()

	req, err := request.New(conn)
	if err != nil {
		return fmt.Errorf("invalid request: %w", err)
	}

	res := pathRouter.Handle(req)
	if err := res.Write(conn); err != nil {
		return fmt.Errorf("failed to write response to the http connection: %w", err)
	}

	return nil
}
