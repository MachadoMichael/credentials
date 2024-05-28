package encrypt

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(payload interface{}) (string, error) {

	tt1 := time.Hour * 24
	secretJWTKey := os.Getenv("JWT_SECRET")
	if secretJWTKey == "" {
		log.Fatal("Error on read secret.")
		return "", fmt.Errorf("Cannot read secret on enviroment variables.")
	}

	token := jwt.New(jwt.SigningMethodES256)

	now := time.Now().UTC()
	claim := token.Claims.(jwt.MapClaims)

	claim["sub"] = payload
	claim["exp"] = now.Add(tt1).Unix()
	claim["iat"] = now.Unix()
	claim["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWTKey))

	if err != nil {
		return "", fmt.Errorf("generation JWT Token failed: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["sub"], nil
}
