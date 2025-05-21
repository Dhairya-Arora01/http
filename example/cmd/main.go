package main

import (
	"fmt"

	"github.com/Dhairya-Arora01/http"
	"github.com/Dhairya-Arora01/http/pkg/log"
	"github.com/Dhairya-Arora01/http/pkg/method"
	"github.com/Dhairya-Arora01/http/pkg/request"
	"github.com/Dhairya-Arora01/http/pkg/response"
	"github.com/Dhairya-Arora01/http/pkg/router"
	"github.com/Dhairya-Arora01/http/pkg/status"
)

func main() {
	logger, err := log.New(log.DebugLevel)
	if err != nil {
		fmt.Printf("failed to initialize logger: %v\n", err)
		return
	}

	pathRouter := router.New()
	pathRouter.RegisterHandler("/coffee", []method.Method{method.GET}, handleGetCoffee)

	if err := http.Start(logger, 8080, pathRouter); err != nil {
		fmt.Printf("encountered an error while listening: %v\n", err)
	}
}

func handleGetCoffee(*request.Request) (*response.Response, error) {
	return response.New(status.OK, nil, nil), nil
}
