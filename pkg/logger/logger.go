package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/exp/slog"
)

// Logger represents a logger.
type Logger struct {
	slogger *slog.Logger
	file    *os.File
}

var LoginLogger *Logger
var ErrorLogger *Logger

func NewLogger(file *os.File, level slog.Level) (*Logger, error) {
	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{Level: level})
	slogger := slog.New(handler)
	return &Logger{slogger: slogger, file: file}, nil
}

func (l *Logger) Write(level slog.Level, message string) {
	ctx := context.Background()
	l.slogger.Log(ctx, level, message, slog.Time("timespam", time.Now()))
}

func (l *Logger) Close() error {
	return l.file.Close()
}

func InitLoggers() {

	loginFile, err := startingFile("login.log")
	if err != nil {
		log.Fatalf("Error on start login.log, error: %v", err)
	}

	errorFile, err := startingFile("error.log")
	if err != nil {
		log.Fatalf("Error on start error.log, error: %v", err)
	}

	loginLogger, err := NewLogger(loginFile, slog.LevelInfo)
	if err != nil {
		log.Fatalf("Error on start loginLogger, error: %v", err)
	}

	errorLogger, err := NewLogger(errorFile, slog.LevelError)
	if err != nil {

		log.Fatalf("Error on start errorLogger, error: %v", err)
	}

	LoginLogger = loginLogger
	ErrorLogger = errorLogger

}

func startingFile(fileName string) (*os.File, error) {
	dir := "/Users/michael/Projects/credentials/log/"

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return nil, err
		}
	}

	filePath := dir + fileName
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalln("Error creating file:", err)
			return nil, err
		}
		return file, nil
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}
