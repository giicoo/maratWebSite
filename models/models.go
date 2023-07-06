package models

import (
	"github.com/golang-jwt/jwt"
)

// for user
type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Datatime string `json:"datatime"`
}

type Claims struct {
	jwt.StandardClaims
	Login string `json:"login"`
}

// for words
type Word struct {
	Word      string `json:"word"`
	Translate string `json:"translate"`
	Datatime  string `json:"datatime"`
}

type CheckTestWord struct {
	Word  *Word  `json:"word"`
	Check bool   `json:"check"`
	Right string `json:"right"`
}

// for tests
type ElemTest struct {
	Words []*Word `json:"words"`
	Right int     `json:"right"`
}

type Test struct {
	Name         string        `json:"name"`
	Words        []*Word       `json:"words"`
	UsersResults []*UserResult `json:"users_results"`
	Datatime     string        `json:"datatime"`
}

type UserResult struct {
	Login    string           `json:"login"`
	Percent  int              `json:"percent"`
	Res      []*CheckTestWord `json:"res"`
	Datatime string           `json:"datatime"`
}
