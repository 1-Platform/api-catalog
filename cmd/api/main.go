package main

import (
	"github.com/1-platform/api-catalog/internal/api"
	"github.com/1-platform/api-catalog/internal/auth"
	"github.com/1-platform/api-catalog/internal/pkg/config"
	"github.com/1-platform/api-catalog/internal/pkg/datastore"
	"github.com/1-platform/api-catalog/internal/pkg/smtp"
	"github.com/1-platform/api-catalog/internal/server/http"
	"github.com/1-platform/api-catalog/internal/teams"
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

	// smtp setup
	smtpService := smtp.NewSmtpService(&smtp.Config{
		From:     cfg.Smtp.From,
		Host:     cfg.Smtp.Host,
		Port:     cfg.Smtp.Port,
		UserName: cfg.Smtp.UserName,
		Password: cfg.Smtp.Password,
	})

	authStore := auth.NewStore(db)
	auth := auth.New(authStore, cfg)

	tmStore := teams.NewStore(db)
	tm := teams.New(tmStore, smtpService, cfg)

	apiService, err := api.New(version, tm)
	if err != nil {
		logger.Fatal(err)
	}

	srvCfg := &http.Config{Port: cfg.Port, Host: cfg.Host}
	httpServer, err := http.New(srvCfg, logger, auth, apiService)
	if err != nil {
		logger.Fatal(err)
	}

	if err := httpServer.Listen(); err != nil {
		logger.Fatal(err)
	}
}
