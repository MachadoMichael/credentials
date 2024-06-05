package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger represents a logger.
type Logger struct {
	filename string
	file     *os.File
}

var LoginLogger *Logger
var ErrorLogger *Logger

func NewLogger(file *os.File) (*Logger, error) {
	// file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	return nil, err
	// }
	return &Logger{filename: file.Name(), file: file}, nil
}

// Write writes a log message to the logger.
func (l *Logger) Write(level string, message string) {
	timestamp := time.Now().Format(time.RFC3339)
	logMessage := fmt.Sprintf("[%s] %s: %s\n", timestamp, level, message)
	_, err := l.file.WriteString(logMessage)
	if err != nil {
		log.Println(err)
	}

	l.Close()
}

// Close closes the logger.
func (l *Logger) Close() error {
	return l.file.Close()
}

// InitLoggers initializes the loggers.
func InitLoggers() error {

	loginFile, err := startingFile("login.log")
	if err != nil {
		return err
	}

	errorFile, err := startingFile("error.log")
	if err != nil {
		return err
	}

	loginLogger, err := NewLogger(loginFile)
	if err != nil {
		return err
	}

	errorLogger, err := NewLogger(errorFile)
	if err != nil {
		return err
	}

	LoginLogger = loginLogger
	ErrorLogger = errorLogger

	return nil
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

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}
