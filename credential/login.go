package credential

import (
	"encoding/json"
	"net/http"

	"github.com/MachadoMichael/credentials/model"
	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"golang.org/x/exp/slog"
)

func (c *credentialHandler) Login(w http.ResponseWriter, r *http.Request) {
	credential := model.Credential{}
	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	credentialPassword, err := c.Repo.ReadOne(credential.Email)
	if credentialPassword == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		c.ErrorLogger.Write(slog.LevelError, "Credential not found")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	err = encrypt.VerifyPassword(credentialPassword, credential.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	token, err := encrypt.GenerateToken(credential.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	c.AccessLogger.Write(slog.LevelInfo, "sucessful login attempt, email: "+credential.Email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}
