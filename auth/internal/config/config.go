package config

import "github.com/fin/tools/pkg/configcore"

// Config конфигурация сервиса авторизации
type Config struct {
	AppName               string `envconfig:"APP_NAME"`
	AppDebug              bool   `envconfig:"APP_DEBUG"`
	CatalogServerConfig   configcore.ServerConfig
	DatabaseURI           string `envconfig:"DATABASE_URI"`
	LogLevel              string `envconfig:"LOG_LEVEL" default:"info"`
	ObservabilitySettings configcore.Observer
}
