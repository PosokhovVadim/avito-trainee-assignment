package main

import (
	"avito/config"
	"avito/internal/app"
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

	app := app.NewBannerApp(log, cfg.HTTPServer.Port, cfg.StoragePath)
	app.Run()
	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
