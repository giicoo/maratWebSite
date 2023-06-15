package models

import "github.com/golang-jwt/jwt"

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:username`
}
