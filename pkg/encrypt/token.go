package encrypt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(payload string) (string, error) {
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

//
// func ValidateToken(token string) (interface{}, error) {
//
// 	var signingKey = []byte(os.Getenv("JWT_SECRET"))
// 	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
// 		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
// 		}
//
// 		return []byte(signingKey), nil
// 	})
//
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid token %w", err)
// 	}
//
// 	claims, ok := tok.Claims.(jwt.MapClaims)
// 	if !ok || !tok.Valid {
// 		return nil, fmt.Errorf("invalid token claim")
// 	}
//
// 	return claims["sub"], nil
// }
