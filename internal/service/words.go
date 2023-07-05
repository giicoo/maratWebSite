package service

import (
	"time"

	"github.com/giicoo/maratWebSite/internal/repository"
	"github.com/giicoo/maratWebSite/models"
)

type WordsServices interface {
	AddWord(w models.WordDB) error
	GetWord() ([]*models.WordDB, error)
}

type WordsService struct {
	repo repository.Repo
}

func (s *WordsService) AddWord(w models.WordDB) error {
	w.Datatime = time.Now().Format(time.ANSIC)
	return s.repo.AddWord(w)
}

func (s *WordsService) GetWord() ([]*models.WordDB, error) {
	return s.repo.GetWords()
}
