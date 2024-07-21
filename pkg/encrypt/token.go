package encrypt

import (
	"fmt"
	"time"

	"github.com/MachadoMichael/credentials/infra"
	"github.com/golang-jwt/jwt"
)

var mySigningKey []byte

func GenerateToken(payload string) (string, error) {
	mySigningKey = []byte(infra.Config.JwtSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	} else {
		return false, err
	}
}
