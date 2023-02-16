package main

import (
	"log"

	"github.com/1-platform/api-catalog/internal/api"
	"github.com/1-platform/api-catalog/internal/server/http"
)

var version = "development"

func main() {
	apiService, err := api.New(version)
	if err != nil {
		log.Fatal(err)
	}

	httpServer, err := http.New(&http.Config{Port: "8000"}, apiService)
	if err != nil {
		log.Fatal(err)
	}

	if err := httpServer.Listen(); err != nil {
		log.Fatal(err)
	}
}
