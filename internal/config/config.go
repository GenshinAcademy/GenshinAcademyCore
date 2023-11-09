package config

import (
	"github.com/Jagerente/gocfg"
)

type Config struct {
	AppConfig
	RouterConfig
	DbConfig
	AssetsConfig
	LoggerConfig
}

func New() (*Config, error) {
	var cfg = new(Config)

	if err := gocfg.NewDefault().Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

type AppConfig struct {
	SecretKey string `env:"SECRET_KEY"`
}

type RouterConfig struct {
	GinMode    string `env:"GIN_MODE"`
	ServerPort uint16 `env:"SERVER_PORT"`
}

type DbConfig struct {
	DBHost         string `env:"POSTGRES_HOST"`
	DBUserName     string `env:"POSTGRES_USER"`
	DBUserPassword string `env:"POSTGRES_PASSWORD"`
	DBName         string `env:"POSTGRES_DB"`
	DBPort         uint16 `env:"POSTGRES_PORT"`
}

type AssetsConfig struct {
	AssetsPath string `env:"ASSETS_PATH"`
	AssetsHost string `env:"ASSETS_HOST"`
}

type LoggerConfig struct {
	LogLevel     int  `env:"LOG_LEVEL"`
	ReportCaller bool `env:"REPORT_CALLER"`
	LogFormatter int  `env:"LOG_FORMATTER"`
}
