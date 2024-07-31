package credentials

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MachadoMichael/credentials/dto"
	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/schema"
	"golang.org/x/exp/slog"
)

func (s *Service) Update(w http.ResponseWriter, r *http.Request) {
	if !s.IsValidToken(w, r) {
		return
	}

	request := dto.UpdatePasswordRequestDTO{}
	credBackup := schema.Credentials{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		s.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	credentialPassword, err := s.Repo.ReadOne(request.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		s.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	err = encrypt.VerifyPassword(request.OldPassword, credentialPassword)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		s.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	credBackup.Email = request.Email
	credBackup.Password = request.OldPassword

	rows, err := s.Repo.Delete(request.Email)
	if err != nil {
		s.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error() + "rows affcteds: " + strconv.FormatInt(rows, 10))
		return
	}

	err = s.Repo.Create(schema.Credentials{Email: request.Email, Password: request.NewPassword})
	if err != nil {
		backErr := s.Repo.Create(credBackup)
		if backErr != nil {
			s.ErrorLogger.Write(slog.LevelError, backErr.Error())
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		s.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	s.AccessLogger.Write(slog.LevelInfo, "Successful password update, email: "+request.Email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Password update successfully")

}
