package infra

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigData struct {
	JwtSecret   string
	LogFilePath string
}

var Config *ConfigData

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Config = &ConfigData{
		JwtSecret:   os.Getenv("JWT_SECRET"),
		LogFilePath: os.Getenv("LOG_FILE_PATH"),
	}

}
