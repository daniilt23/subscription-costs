package logger

import (
	"os"

	"go.uber.org/zap"
)

func InitLogger() *zap.Logger {
	logger := zap.Must(zap.NewDevelopment())
	if os.Getenv("ENV") == "PROD" {
		return zap.Must(zap.NewProduction())
	}

	return logger
}