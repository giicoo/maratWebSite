package service

import "github.com/giicoo/maratWebSite/models"

func (s *Services) AddWord(w models.Word) error {
	return s.repo.AddWord(w.Word, w.Translate)
}

func (s *Services) GetWord() ([]models.Word, error) {
	return s.repo.GetWords()
}
