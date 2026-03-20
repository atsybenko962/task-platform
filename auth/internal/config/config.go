package config

import (
	"github.com/fin/tools/pkg/configcore"
	"time"
)

type Config struct {
	Env            string        `yaml:"env" env-default:"local"`
	StoragePath    string        `yaml:"storage_path" env-required:"true"`
	GRPC           GRPCConfig    `yaml:"grpc"`
	MigrationsPath string
	TokenTTL       time.Duration `yaml:"token_ttl" env-default:"1h"`
	LogLevel       string        `yaml:"log_level" env-default:"info"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad(path string) (*Config, error) {
	var cfg Config
	if err := configcore.MustLoad(&cfg, path); err != nil {
		return nil, err
	}
	return &cfg, nil
}
