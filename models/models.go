package models

import (
	"github.com/golang-jwt/jwt"
)

type UserDB struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type WordDB struct {
	Word      string `json:"word"`
	Translate string `json:"translate"`
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

type TestWord struct {
	Word  *WordDB `json:"word"`
	Check bool    `json:"check"`
	Right string  `json:"right"`
}

type WorkTest struct {
	Words []*WordDB `json:"words"`
	Right int       `json:"right"`
}
