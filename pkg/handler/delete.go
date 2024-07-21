package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"golang.org/x/exp/slog"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	var email string
	err := json.NewDecoder(r.Body).Decode(&email)
	if !isValidToken(w, r) {
		return
	}

	cred, errRead := database.CredentialRepo.ReadOne(email)
	if errRead != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		logger.ErrorLogger.Write(slog.LevelError, errRead.Error())
		return
	}

	if cred == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("There isn't register for this email.")
		logger.ErrorLogger.Write(slog.LevelError, "There isn't register for this email.")
		return
	}

	rows, err := database.CredentialRepo.Delete(email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	logger.AccessLogger.Write(slog.LevelInfo, "Credential deleted successfully, email: "+email+"rows: "+strconv.FormatInt(rows, 10))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Credential deleted successfully, email: " + email)
}
