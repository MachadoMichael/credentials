package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang.org/x/exp/slog"
)

func (c *credentialHandler) Delete(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")
	if !c.IsValidToken(w, r) {
		return
	}

	cred, err := s.Repo.ReadOne(email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	if cred == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("There isn't register for this email.")
		c.ErrorLogger.Write(slog.LevelError, "There isn't register for this email.")
		return
	}

	rows, err := c.Repo.Delete(email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	c.AccessLogger.Write(slog.LevelInfo, "Credential deleted successfully, email: "+email+"rows: "+strconv.FormatInt(rows, 10))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Credential deleted successfully, email: " + email)

}
