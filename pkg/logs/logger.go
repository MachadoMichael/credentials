package logs

import (
	"log/slog"
	"sync"
)

var (
	loginLogger *slog.Logger
	errorLogger *slog.Logger
	once        sync.Once
)
