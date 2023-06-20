package service

import (
	"github.com/giicoo/maratWebSite/models"
)

func (s *Services) AddWord(w models.WordDB) error {
	return s.repo.AddWord(w)
}

func (s *Services) GetWord() ([]*models.WordDB, error) {
	return s.repo.GetWords()
}

func (s *Services) CheckTest(words []*models.WordDB) ([]*models.TestWord, error) {
	answers, err := s.repo.GetWordsByNames(words)
	if err != nil {
		return nil, err
	}
	answersMap := map[string]string{}
	for _, item := range answers {
		answersMap[item.Word] = item.Translate
	}

	test_answers := []*models.TestWord{}

	for _, item := range words {
		test_elm := models.TestWord{
			Word:  item,
			Check: item.Translate == answersMap[item.Word],
			Right: answersMap[item.Word],
		}
		test_answers = append(test_answers, &test_elm)
	}

	return test_answers, nil
}
