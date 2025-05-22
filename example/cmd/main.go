package main

import (
	"bytes"
	"encoding/json"
	"errors"
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
	pathRouter.RegisterHandler("/brokenmachine", []method.Method{method.GET}, handleBrokenCoffeeMachine)

	if err := http.Start(logger, 8080, pathRouter); err != nil {
		fmt.Printf("encountered an error while listening: %v\n", err)
	}
}

type Coffee struct {
	Message     string `json:"message"`
	Temperature int    `json:"temperature"`
}

func handleGetCoffee(*request.Request) (*response.Response, error) {
	coffee := Coffee{
		Message:     "Here is your coffee",
		Temperature: 63,
	}

	coffeeData, err := json.Marshal(coffee)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal coffee")
	}

	return response.New(status.OK, response.ContentTypeJSON, nil, bytes.NewReader(coffeeData)), nil
}

func handleBrokenCoffeeMachine(*request.Request) (*response.Response, error) {
	return nil, errors.New("broken coffee machine")
}
