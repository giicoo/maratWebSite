package service

import (
	"math/rand"
	"time"

	tools_service "github.com/giicoo/maratWebSite/internal/service/tools"
	"github.com/giicoo/maratWebSite/models"
)

func (s *Services) AddWord(w models.WordDB) error {
	return s.repo.AddWord(w)
}

func (s *Services) GetWord() ([]*models.WordDB, error) {
	return s.repo.GetWords()
}

func (s *Services) GetWordsForTest() ([]*models.WorkTest, error) {
	tests := []*models.WorkTest{}

	words, err := s.repo.GetWords()
	if err != nil {
		return tests, err
	}

	for i, item := range words {
		test := models.WorkTest{}
		test.Words = make([]*models.WordDB, 4)
		rand.Seed(time.Now().UnixNano())
		root_i := rand.Intn(4)
		test.Words = tools_service.RandomWords(words, i)
		test.Words = tools_service.Insert(test.Words, root_i, item)
		test.Right = root_i

		tests = append(tests, &test)
	}
	return tests, nil
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
