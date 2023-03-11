package configuration

import (
	"go.uber.org/zap"
)

func GetLogger() *zap.Logger {
	logger, _ := zap.NewProduction()

	if ENV.LogLevel == 1 {
		logger, _ = zap.NewDevelopment()
	}

	return logger
}
