package configuration

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         uint16 `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	LogLevel       uint16 `mapstructure:"LOG_LEVEL"`
	GinMode        string `mapstructure:"GIN_MODE"`
	SecretKey      string `mapstructure:"SECRET_KEY"`
	AssetsPath     string `mapstructure:"ASSETS_PATH"`
	AssetsHost     string `mapstructure:"ASSETS_HOST"`
	AssetsFormat   string `mapstructure:"ASSETS_FORMAT"`
}

var (
	ENV Config
)

func loadENV() error {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		ENV.DBHost = viper.GetString("POSTGRES_HOST")
		ENV.DBUserName = viper.GetString("POSTGRES_USER")
		ENV.DBUserPassword = viper.GetString("POSTGRES_PASSWORD")
		ENV.DBName = viper.GetString("POSTGRES_DB")
		ENV.DBPort = viper.GetUint16("POSTGRES_PORT")
		ENV.ServerPort = viper.GetString("SERVER_PORT")
		ENV.LogLevel = viper.GetUint16("LOG_LEVEL")
		ENV.GinMode = viper.GetString("GIN_MODE")
		ENV.SecretKey = viper.GetString("SECRET_KEY")
	} else {
		if err := viper.Unmarshal(&ENV); err != nil {
			return err
		}
	}

	return nil
}

func Init() error {
	if err := loadENV(); err != nil {
		return err
	}

	return nil
}
