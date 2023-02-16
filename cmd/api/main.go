package main

import (
	"github.com/1-platform/api-catalog/internal/api"
	"github.com/1-platform/api-catalog/internal/server/http"
	"github.com/1-platform/api-catalog/pkg/logger"
)

var version = "development"

func main() {
	logger, err := logger.New(version != "development")
	if err != nil {
		logger.Fatal(err)
	}

	apiService, err := api.New(version)
	if err != nil {
		logger.Fatal(err)
	}

	httpServer, err := http.New(&http.Config{Port: "8000"},
		logger,
		apiService,
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := httpServer.Listen(); err != nil {
		logger.Fatal(err)
	}
}
