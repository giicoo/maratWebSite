package service

import (
	"time"

	"github.com/giicoo/maratWebSite/internal/repository"
	"github.com/giicoo/maratWebSite/models"
)

type WordsFuncs interface {
	AddWord(w models.Word) error
	GetWord() ([]*models.Word, error)
	GetWordsByNames(words []*models.Word) ([]*models.Word, error)
}

type WordsService struct {
	repo repository.Repo
}

func (s *WordsService) AddWord(w models.Word) error {
	// set data time
	w.Datatime = time.Now().Format(time.ANSIC)

	// add word
	return s.repo.AddWord(w)
}

func (s *WordsService) GetWord() ([]*models.Word, error) {
	// get words
	return s.repo.GetWords()
}

func (s *WordsService) GetWordsByNames(words []*models.Word) ([]*models.Word, error) {
	return s.repo.GetWordsByNames(words)
}
