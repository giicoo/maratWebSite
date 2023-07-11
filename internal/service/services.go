package service

import (
	"github.com/giicoo/maratWebSite/internal/repository"
	hashFunc "github.com/giicoo/maratWebSite/pkg/hash_password"
)

type Services struct {
	AuthServices       AuthFuncs
	WordsServices      WordsFuncs
	TestServices       TestFuncs
	StatisticsServices StatisticsFunc
}

func NewServices(repo repository.Repo, hash hashFunc.HashTools) *Services {
	return &Services{
		AuthServices:       &AuthService{repo: repo, hashTools: hash},
		WordsServices:      &WordsService{repo: repo},
		TestServices:       &TestService{repo: repo},
		StatisticsServices: &StatisticsService{path: "./files/stat.xlsx"},
	}
}
