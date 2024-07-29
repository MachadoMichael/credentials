package logger

import (
	"log"

	"golang.org/x/exp/slog"
)

func Init() (*Logger, *Logger, error) {

	accessFile, err := loadingFile("access.log")
	if err != nil {
		log.Fatalf("Error on start access.log, error: %v", err)
	}

	errorFile, err := loadingFile("error.log")
	if err != nil {
		log.Fatalf("Error on start error.log, error: %v", err)
	}

	accessLogger, err := NewLogger(accessFile, slog.LevelInfo)
	if err != nil {
		log.Fatalf("Error on start accessLogger, error: %v", err)
	}

	errorLogger, err := NewLogger(errorFile, slog.LevelError)
	if err != nil {

		log.Fatalf("Error on start errorLogger, error: %v", err)
	}

	return accessLogger, errorLogger, nil
}
