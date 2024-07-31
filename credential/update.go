package credential

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MachadoMichael/credentials/dto"
	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/schema"
	"golang.org/x/exp/slog"
)

func (c *CredentialHandler) Update(w http.ResponseWriter, r *http.Request) {
	if !c.IsValidToken(w, r) {
		return
	}

	request := dto.UpdatePasswordRequestDTO{}
	credBackup := schema.Credentials{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	credentialPassword, err := c.Repo.ReadOne(request.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	err = encrypt.VerifyPassword(request.OldPassword, credentialPassword)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	credBackup.Email = request.Email
	credBackup.Password = request.OldPassword

	rows, err := c.Repo.Delete(request.Email)
	if err != nil {
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error() + "rows affcteds: " + strconv.FormatInt(rows, 10))
		return
	}

	err = c.Repo.Create(schema.Credentials{Email: request.Email, Password: request.NewPassword})
	if err != nil {
		err = c.Repo.Create(credBackup)
		if err != nil {
			c.ErrorLogger.Write(slog.LevelError, err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		c.ErrorLogger.Write(slog.LevelError, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	c.AccessLogger.Write(slog.LevelInfo, "Successful password update, email: "+request.Email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Password update successfully")

}
