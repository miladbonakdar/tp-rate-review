package logger

import (
	"log"

	"go.uber.org/zap"
)

func Log() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println(err)
	}
	return logger
}
