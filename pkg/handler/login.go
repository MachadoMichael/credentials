package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/MachadoMichael/credentials/schema"
	"golang.org/x/exp/slog"
)

func Login(w http.ResponseWriter, r *http.Request) {
	credential := schema.Credentials{}
	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	credentialPassword, err := database.CredentialRepo.ReadOne(credential.Email)
	if credentialPassword == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		logger.ErrorLogger.Write(slog.LevelError, "Credential not found")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	err = encrypt.VerifyPassword(credentialPassword, credential.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	token, err := encrypt.GenerateToken(credential.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	logger.AccessLogger.Write(slog.LevelInfo, "sucessful login attempt, email: "+credential.Email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
