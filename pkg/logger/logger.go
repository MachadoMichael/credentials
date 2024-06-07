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

var AccessLogger *Logger
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

	accessFile, err := startingFile("access.log")
	if err != nil {
		log.Fatalf("Error on start access.log, error: %v", err)
	}

	errorFile, err := startingFile("error.log")
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

	AccessLogger = accessLogger
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
