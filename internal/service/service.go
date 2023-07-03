package service

import (
	"github.com/giicoo/maratWebSite/internal/repository"
	hashFunc "github.com/giicoo/maratWebSite/pkg/hash_password"
)

type Services struct {
	AuthServices  AuthServices
	WordsServices WordsServices
	TestServices  TestServices
}

func NewServices(repo repository.Repo, hash hashFunc.HashTools) *Services {
	return &Services{
		AuthServices:  &AuthService{repo: repo, hashTools: hash},
		WordsServices: &WordsService{repo: repo},
		TestServices:  &TestService{repo: repo},
	}
}
