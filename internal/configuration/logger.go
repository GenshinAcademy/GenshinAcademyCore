package configuration

import (
	"go.uber.org/zap"
)

func setupLogger() {
	Logger, _ := zap.NewProduction()

	if ENV.LogLevel == 1 {
		Logger, _ = zap.NewDevelopment()
	}

	zap.ReplaceGlobals(Logger)
}
