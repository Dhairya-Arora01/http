// Package http provides the HTTP server implementation.
package http

import (
	"bufio"
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
		if err := handleConnection(conn); err != nil {
			log.Error(err, "Failed to handle connection")
		}
	}
}

func handleConnection(conn net.Conn) error {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Read start-line for the request.
	// <method> <request-target> <protocol>.
	startLine, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("invalid request: without start line of the form '<method> <request-target> - <protocol>': %w", err)
	}

	fmt.Println("startLine", startLine)

	// Read headers.
	for {
		header, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("invalid request: failed to read headers: %w", err)
		}

		if header == "\r\n" {
			break
		}

		fmt.Println("header", header)
	}

	return nil
}
