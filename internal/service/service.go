package service

import (
	"github.com/giicoo/maratWebSite/internal/repository"
)

type Services struct {
	repo repository.Repo
}

func NewServices(repo repository.Repo) *Services {
	return &Services{
		repo: repo,
	}
}
