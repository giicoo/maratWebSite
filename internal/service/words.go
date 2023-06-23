package service

import (
	"math/rand"
	"time"

	"github.com/giicoo/maratWebSite/internal/repository"
	tools_service "github.com/giicoo/maratWebSite/internal/service/tools"
	"github.com/giicoo/maratWebSite/models"
)

type WordsServices interface {
	AddWord(w models.WordDB) error
	GetWord() ([]*models.WordDB, error)
	GetWordsForTest() ([]*models.WorkTest, error)
	CheckTest(words []*models.WordDB) ([]*models.TestWord, error)
}

type WordsService struct {
	repo repository.Repo
}

func (s *WordsService) AddWord(w models.WordDB) error {
	return s.repo.AddWord(w)
}

func (s *WordsService) GetWord() ([]*models.WordDB, error) {
	return s.repo.GetWords()
}

func (s *WordsService) GetWordsForTest() ([]*models.WorkTest, error) {
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

func (s *WordsService) CheckTest(words []*models.WordDB) ([]*models.TestWord, error) {
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
