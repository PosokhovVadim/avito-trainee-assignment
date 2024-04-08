package main

import (
	"avito/config"
	"avito/pkg/logger"
	"fmt"
)

func run() error {
	//todo: init docker - DONE but without db

	cfg, err := config.New("config/local.yml")
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	log, err := logger.SetupLogger(cfg)
	if err != nil {
		return fmt.Errorf("failed to setup logger: %w", err)
	}

	log.Info("Logger setup successfully")

	//todo: app init banner service
	//todo: to app: 1) init logger 2) init db 3) init http server

	//todo: run http server

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
