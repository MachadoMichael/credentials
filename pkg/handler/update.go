package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MachadoMichael/credentials/dto"
	"github.com/MachadoMichael/credentials/infra/database"
	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/pkg/logger"
	"github.com/MachadoMichael/credentials/schema"
	"golang.org/x/exp/slog"
)

func Update(w http.ResponseWriter, r *http.Request) {

	if !isValidToken(w, r) {
		return
	}

	request := dto.UpdatePasswordRequest{}
	credBackup := schema.Credentials{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	credentialPassword, err := database.CredentialRepo.ReadOne(request.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	err = encrypt.VerifyPassword(request.OldPassword, credentialPassword)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	credBackup.Email = request.Email
	credBackup.Password = request.OldPassword

	rows, err := database.CredentialRepo.Delete(request.Email)
	if err != nil {
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error() + "rows affcteds: " + strconv.FormatInt(rows, 10))
		return
	}

	err = database.CredentialRepo.Create(schema.Credentials{Email: request.Email, Password: request.NewPassword})
	if err != nil {
		backErr := database.CredentialRepo.Create(credBackup)
		if backErr != nil {
			logger.ErrorLogger.Write(slog.LevelError, backErr.Error())
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	logger.AccessLogger.Write(slog.LevelInfo, "Successful password update, email: "+request.Email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Password update successfully")

}
