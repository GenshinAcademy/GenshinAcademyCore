package configuration

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         uint16 `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	LogLevel       byte   `mapstructure:"LOG_LEVEL"`
}

var (
	ENV    Config
	Logger *zap.Logger
)

func loadENV() error {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&ENV)
	if err != nil {
		return err
	}

	return nil
}

func Init() error {
	err := loadENV()
	if err != nil {
		return err
	}

	setupLogger()

	return nil
}
