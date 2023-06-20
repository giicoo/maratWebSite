package models

import (
	"github.com/golang-jwt/jwt"
)

type UserDB struct {
	Login    string `bson:"login"`
	Password string `bson:"password"`
}

type WordDB struct {
	Word      string `bson:"word"`
	Translate string `bson:"translate"`
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

type TestWord struct {
	Word  *WordDB
	Check bool
	Right string
}
