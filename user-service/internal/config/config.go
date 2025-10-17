package config

import "github.com/task_platform/tools/configcore"

type Config struct {
	AppName               string `envconfig:"APP_NAME"`
	AppDebug              bool   `envconfig:"APP_DEBUG"`
	UserServerConfig      configcore.ServerConfig
	DatabaseURI           string `envconfig:"DATABASE_URI"`
	LogLevel              string `envconfig:"LOG_LEVEL" default:"info"`
	ObservabilitySettings configcore.Observer
}
