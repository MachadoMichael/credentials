package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/MachadoMichael/credentials/dto"
	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/schema"
	"golang.org/x/exp/slog"
)

func (s *Service) Update(w http.ResponseWriter, r *http.Request) error {

	if !s.IsValidToken(w, r) {
		return errors.New("invalid token.")
	}

	request := dto.UpdatePasswordRequestDTO{}
	credBackup := schema.Credentials{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		s.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return err
	}

	credentialPassword, err := s.Repo.ReadOne(request.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		s.ErrorLogger.Write(slog.LevelError, err.Error())
		return err
	}

	err = encrypt.VerifyPassword(request.OldPassword, credentialPassword)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		s.ErrorLogger.Write(slog.LevelError, err.Error())
		return err
	}

	credBackup.Email = request.Email
	credBackup.Password = request.OldPassword

	rows, err := s.Repo.Delete(request.Email)
	if err != nil {
		s.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error() + "rows affcteds: " + strconv.FormatInt(rows, 10))
		return err
	}

	err = s.Repo.Create(schema.Credentials{Email: request.Email, Password: request.NewPassword})
	if err != nil {
		backErr := s.Repo.Create(credBackup)
		if backErr != nil {
			s.ErrorLogger.Write(slog.LevelError, backErr.Error())
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err.Error())
			return err
		}

		s.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return err
	}

	s.AccessLogger.Write(slog.LevelInfo, "Successful password update, email: "+request.Email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Password update successfully")
	return nil

}
