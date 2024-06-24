package infra

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigData struct {
	DbAddress   string
	DbPassword  string
	DbName      int
	JwtSecret   string
	LogFilePath string
}

var Config *ConfigData

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	envDbName := os.Getenv("DATABASE_NAME")
	if envDbName == "" {
		log.Fatal("Error on read database name.")
	}

	dbName, err := strconv.Atoi(envDbName)
	if err != nil {
		log.Fatal(err)
		panic("Cannot read envDbName")
	}

	Config = &ConfigData{
		DbAddress:   os.Getenv("DATABASE_ADDRESS"),
		DbPassword:  os.Getenv("DATABASE_PASSWORD"),
		DbName:      dbName,
		JwtSecret:   os.Getenv("JWT_SECRET"),
		LogFilePath: os.Getenv("LOG_FILE_PATH"),
	}

}
