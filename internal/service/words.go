package service

import "github.com/giicoo/maratWebSite/models"

func (s *Services) AddWord(w models.WordDB) error {
	return s.repo.AddWord(w)
}

func (s *Services) GetWord() ([]*models.WordDB, error) {
	return s.repo.GetWords()
}
