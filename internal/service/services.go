package service

import (
	"github.com/giicoo/maratWebSite/configs"
	"github.com/giicoo/maratWebSite/internal/repository"
	hashFunc "github.com/giicoo/maratWebSite/pkg/hash_password"
)

type Services struct {
	AuthServices       AuthFuncs
	WordsServices      WordsFuncs
	TestServices       TestFuncs
	StatisticsServices StatisticsFunc
}

func NewServices(repo repository.Repo, hash hashFunc.HashTools, cfg *configs.Config) *Services {
	return &Services{
		AuthServices:       &AuthService{repo: repo, hashTools: hash, cfg: cfg},
		WordsServices:      &WordsService{repo: repo, cfg: cfg},
		TestServices:       &TestService{repo: repo, cfg: cfg},
		StatisticsServices: &StatisticsService{path: "./files/stat.xlsx", cfg: cfg},
	}
}
