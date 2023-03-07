package main

import (
	"github.com/1-platform/api-catalog/internal/api"
	"github.com/1-platform/api-catalog/internal/auth"
	"github.com/1-platform/api-catalog/internal/pkg/config"
	"github.com/1-platform/api-catalog/internal/pkg/datastore"
	"github.com/1-platform/api-catalog/internal/server/http"
	"github.com/1-platform/api-catalog/pkg/logger"
)

var version = "development"

func main() {
	logger, err := logger.New(version != "development")
	if err != nil {
		logger.Fatal(err)
	}

	cfg, err := config.New(".")
	if err != nil {
		logger.Fatal(err)
	}

	// setup mongodb connection
	db, err := datastore.New(cfg.MongoDbURI, cfg.MongoDbName)
	if err != nil {
		logger.Fatal(err)
	}

	authStore := auth.NewAuthStore(db)
	auth := auth.New(authStore, cfg)

	apiService, err := api.New(version)
	if err != nil {
		logger.Fatal(err)
	}

	httpServer, err := http.New(&http.Config{Port: cfg.Port},
		logger,
		auth,
		apiService,
	)

	if err != nil {
		logger.Fatal(err)
	}

	if err := httpServer.Listen(); err != nil {
		logger.Fatal(err)
	}
}
