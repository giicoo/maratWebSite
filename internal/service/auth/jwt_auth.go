package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/giicoo/maratWebSite/models"
	"github.com/golang-jwt/jwt"
)

func NewJWT(login string) (string, error) {
	// generate jwt token by user login with
	// + ExpiresAt - time for access
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(5 * time.Minute).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
		Login: login,
	})
	return token.SignedString([]byte("user"))
}

func ParseJWT(tk string) (string, error) {
	// parse token
	token, err := jwt.ParseWithClaims(tk, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// func to check method how hash token and send key for this token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("user"), nil
	})
	if err != nil {
		return "", err
	}

	// check valid token
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims.Login, nil
	}

	return "", errors.New("Invalid token")
}
