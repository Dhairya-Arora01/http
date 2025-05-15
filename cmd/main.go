package main

import (
	"fmt"
	"log"
	"net"
)

var helloHTML = `<html>
<head>
    <title>Hello Page</title>
</head>
<body>
    <h1>Hello, Dhairya!</h1>
    <p>This is a basic HTML response from a Go server.</p>
</body>
</html>`

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to bind to port 8080: %s", err.Error())
	}

	defer listener.Close()

	addr := listener.Addr()
	fmt.Printf("Listening on address %s\n", addr.String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %s\n", err.Error())
			continue
		}

		if err := handleConnection(conn); err != nil {
			fmt.Printf("Failed to handle connection: %s\n", err.Error())
			conn.Close()
			continue
		}

		conn.Close()
	}
}

func handleConnection(conn net.Conn) error {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("failed to read data from connection: %w", err)
	}

	body := []byte(helloHTML)
	headers := fmt.Sprintf("HTTP/1.1 200 OK\r\n"+
		"Content-Type: text/html\r\n"+
		"Content-Length: %d\r\n"+
		"Connection: close\r\n"+
		"\r\n", len(body))

	// Send headers
	if _, err := conn.Write([]byte(headers)); err != nil {
		return fmt.Errorf("failed to write headers: %w", err)
	}

	// Send body
	if _, err := conn.Write(body); err != nil {
		return fmt.Errorf("failed to write body: %w", err)
	}

	return nil
}
