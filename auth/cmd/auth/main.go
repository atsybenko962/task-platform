package main

import (
	"log"
	"os"

	"auth/internal/config"

	"github.com/fin/tools/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg, err := config.MustLoad("config/config_local.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	zapLogger := logger.NewLogger(logger.Config{
		AppName:  "auth",
		LogLevel: cfg.LogLevel,
	}, zapcore.AddSync(os.Stdout))
	defer zapLogger.Sync()

	zapLogger.Info("logger initialized", zap.String("env", cfg.Env))
}
