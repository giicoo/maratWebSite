package models

import (
	"github.com/golang-jwt/jwt"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserDB struct {
	Login    string `bson:"login"`
	Password string `bson:"password"`
}

type Word struct {
	Word      string `json:"word"`
	Translate string `json:"translate"`
}

type WordDB struct {
	Word      string `bson:"word"`
	Translate string `bson:"translate"`
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:username`
}
