// Package http provides the HTTP server implementation.
package http

import (
	"fmt"
	"net"

	"github.com/Dhairya-Arora01/http/pkg/port"
	"github.com/go-logr/logr"
)

// Start listens and accept connections at the specified port.
func Start(log logr.Logger, p int) error {
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

		// TODO: handle the request.

		conn.Close()
	}
}
