package logger

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

type Logger struct {
	slogger *slog.Logger
	file    *os.File
}

func NewLogger(file *os.File, level slog.Level) (*Logger, error) {
	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{Level: level})
	slogger := slog.New(handler)
	return &Logger{slogger: slogger, file: file}, nil
}

func (l *Logger) Write(level slog.Level, message string) {
	ctx := context.Background()
	l.slogger.Log(ctx, level, message)
}

func (l *Logger) Close() error {
	return l.file.Close()
}
