package service

import (
	"crypto/rand"
	"math"
	"math/big"
	"time"

	"github.com/giicoo/maratWebSite/internal/repository"
	"github.com/giicoo/maratWebSite/internal/service/tools"
	"github.com/giicoo/maratWebSite/models"
)

type TestServices interface {
	GetTestByName(name string) (models.Test, error)
	AddTest(words models.Test) error
	GetWordsForTest(name string) ([]*models.WorkTest, error)
	CheckTest(words []*models.WordDB, test_name, username string) ([]*models.TestWord, error)
}

type TestService struct {
	repo repository.Repo
}

func (s *TestService) GetWordsForTest(name string) ([]*models.WorkTest, error) {
	tests := []*models.WorkTest{}

	words, err := s.repo.GetWords()
	if err != nil {
		return tests, err
	}

	testwords, err := s.repo.GetTestByName(name)
	if err != nil {
		return tests, err
	}

	for _, item := range testwords.Words {
		test := models.WorkTest{}
		test.Words = make([]*models.WordDB, 4)
		root_i, _ := rand.Int(rand.Reader, big.NewInt(4))

		test.Words = tools.RandomWords(words, item)
		test.Words = tools.Insert(test.Words, int(root_i.Int64()), item)
		test.Right = int(root_i.Int64())

		tests = append(tests, &test)
	}
	return tests, nil
}

func (s *TestService) CheckTest(words []*models.WordDB, test_name, username string) ([]*models.TestWord, error) {
	answers, err := s.repo.GetWordsByNames(words)
	if err != nil {
		return nil, err
	}
	answersMap := map[string]string{}
	for _, item := range answers {
		answersMap[item.Word] = item.Translate
	}

	test_answers := []*models.TestWord{}

	percent_i := 0
	for _, item := range words {
		test_elm := models.TestWord{
			Word:  item,
			Check: item.Translate == answersMap[item.Word],
			Right: answersMap[item.Word],
		}
		if test_elm.Check {
			percent_i++
		}
		test_answers = append(test_answers, &test_elm)
	}

	percent := math.Round((float64(percent_i) * 100 / float64(len(test_answers))))
	time_res := time.Now().Format(time.ANSIC)

	res := models.UserResult{Login: username, Percent: int(percent), Res: test_answers, Datatime: time_res}
	if err = s.repo.AddUserRes(res, test_name); err != nil {
		return nil, err
	}
	return test_answers, nil
}

func (s *TestService) AddTest(test models.Test) error {
	test.Datatime = time.Now().Format(time.ANSIC)
	return s.repo.AddTest(test)
}

func (s *TestService) GetTestByName(name string) (models.Test, error) {
	return s.repo.GetTestByName(name)
}
