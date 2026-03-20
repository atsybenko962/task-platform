package configcore

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config[T any] struct{}

func MustLoad(cfg any, configPath string) error {
	if configPath == "" {
		configPath = fetchConfigPath()
	}

	if configPath == "" {
		return &ConfigError{Message: "config path is empty"}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &ConfigError{Message: "config file does not exist: " + configPath}
	}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return &ConfigError{Message: "failed to read config: " + err.Error()}
	}

	return nil
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

type ConfigError struct {
	Message string
}

func (e *ConfigError) Error() string {
	return e.Message
}
