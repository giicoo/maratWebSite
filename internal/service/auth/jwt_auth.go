package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/giicoo/maratWebSite/models"
	"github.com/golang-jwt/jwt"
)

func NewJWT(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(5 * time.Minute).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
		Username: login,
	})
	return token.SignedString([]byte("user"))
}

func ParseJWT(tk string) (string, error) {
	token, err := jwt.ParseWithClaims(tk, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("user"), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", errors.New("Invalid Token")
}
