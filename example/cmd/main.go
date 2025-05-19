package main

import (
	"fmt"

	"github.com/Dhairya-Arora01/http"
	"github.com/Dhairya-Arora01/http/pkg/log"
)

func main() {
	logger, err := log.New(log.DebugLevel)
	if err != nil {
		fmt.Printf("failed to initialize logger: %v\n", err)
		return
	}

	if err := http.Start(logger, 8080); err != nil {
		fmt.Printf("encountered an error while listening: %v\n", err)
	}
}
