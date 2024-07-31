package credentials

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MachadoMichael/credentials/pkg/encrypt"
	"golang.org/x/exp/slog"
)

func (c *credentialHandler) IsValidToken(w http.ResponseWriter, r *http.Request) bool {
	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("no token provided")
		c.ErrorLogger.Write(slog.LevelError, "no token provided")
		return false
	}

	strippedTokenStr := strings.TrimPrefix(token, "Bearer ")
	res, err := encrypt.ValidateToken(strippedTokenStr)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("no token provided")
		c.ErrorLogger.Write(slog.LevelError, err.Error())
		return false
	}

	c.AccessLogger.Write(slog.LevelInfo, "Token validate successfully, token "+token)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("token authorized successfully")
	return res

}
