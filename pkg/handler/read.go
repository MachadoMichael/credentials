package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MachadoMichael/credentials/dto"
	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"golang.org/x/exp/slog"
)

func Read(w http.ResponseWriter, r *http.Request) {
	if !isValidToken(w, r) {
		return
	}

	credentials, err := database.CredentialRepo.Read()
	if err != nil {
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	var content []string

	for key := range credentials {
		content = append(content, key)
	}

	layout := "2006-01-02 15:04:05"
	formattedTime := time.Now().Format(layout)

	response := dto.CredentialsRegisteredDTO{
		Lenght:  len(content),
		Content: content,
		ReadAt:  formattedTime,
	}

	logger.AccessLogger.Write(slog.LevelInfo, "Successful to read credentials on database.")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
