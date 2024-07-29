package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang.org/x/exp/slog"
)

func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")
	if !s.IsValidToken(w, r) {
		return
	}

	cred, err := s.Repo.ReadOne(email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		s.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	if cred == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("There isn't register for this email.")
		s.ErrorLogger.Write(slog.LevelError, "There isn't register for this email.")
		return
	}

	rows, err := s.Repo.Delete(email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		s.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	s.AccessLogger.Write(slog.LevelInfo, "Credential deleted successfully, email: "+email+"rows: "+strconv.FormatInt(rows, 10))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Credential deleted successfully, email: " + email)
}
