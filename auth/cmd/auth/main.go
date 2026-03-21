package main

import (
	"auth/internal/config"
	"log"

	"github.com/fin/tools/pkg/logger"
)

func main() {
	cfg, err := config.MustLoad("")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logg := logger.SetupLogger(cfg.Env)

	logg.Info("Lalalal")

}
