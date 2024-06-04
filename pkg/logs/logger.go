package logs

import (
	"log"
	"os"
	"sync"
)

var (
	loginLogger *log.Logger
	errorLogger *log.Logger
	once        sync.Once
)

func startLogFile(filePath string) *log.Logger {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return log.New(file, "", log.LstdFlags|log.Lshortfile)
}

func InitLoggers() {
	once.Do(func() {
		loginLogger = startLogFile("logs/login.log")
		errorLogger = startLogFile("logs/error.log")
	})
}

func LoginLogger() *log.Logger {
	return loginLogger
}

func ErrorLogger() *log.Logger {
	return errorLogger
}
