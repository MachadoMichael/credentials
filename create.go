package credentials

import (
	"encoding/json"
	"net/http"
	"unicode/utf8"

	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"github.com/MachadoMichael/credentials/schema"
	"golang.org/x/exp/slog"
)

func (c *credentialHandler) Create(w http.ResponseWriter, r *http.Request) {
	credential := schema.Credentials{}
	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	cred, err := c.Repo.ReadOne(credential.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	if cred != "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("This email is already in use")
		c.ErrorLogger.Write(slog.LevelError, "This email is already iin use")
		return
	}

	if utf8.RuneCountInString(credential.Password) < 6 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("The password must be longer than 6 characters.")
		c.ErrorLogger.Write(slog.LevelError, "The password must be longer than 6 characters.")
		return
	}

	hash, hashErr := encrypt.HashPassword(credential.Password)
	if hashErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		c.ErrorLogger.Write(slog.LevelError, hashErr.Error())
		return
	}

	credential.Password = hash
	err = c.Repo.Create(credential)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	c.AccessLogger.Write(slog.LevelInfo, "New Credential created successfully, email: "+credential.Email)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Credential created successfully")

}
