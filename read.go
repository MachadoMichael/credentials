package credentials

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MachadoMichael/credentials/dto"
	"golang.org/x/exp/slog"
)

func (s *Service) Read(w http.ResponseWriter, r *http.Request) {
	if !s.IsValidToken(w, r) {
		return
	}

	credentials, err := s.Repo.Read()
	if err != nil {
		s.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	var content []string

	for key := range credentials {
		content = append(content, key)
	}

	layout := "2006-01-02 15:04:05"
	formattedTime := time.Now().Format(layout)

	response := dto.CredentialsRegisteredDTO{
		Lenght:  len(content),
		Content: content,
		ReadAt:  formattedTime,
	}

	s.AccessLogger.Write(slog.LevelInfo, "Successful to read credentials on database.")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
