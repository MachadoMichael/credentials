package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/MachadoMichael/credentials/infra"
)

func loadingFile(fileName string) (*os.File, error) {
	dir := infra.Config.LogFilePath

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

	if err != nil {
		return nil, err
	}
	return file, nil
}
