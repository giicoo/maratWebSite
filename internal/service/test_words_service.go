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

type TestFuncs interface {
	GetTestByName(name string) (models.Test, error)
	GetTests() ([]*models.Test, error)
	AddTest(words models.Test) error
	GetWordsForTest(name string) ([]*models.ElemTest, error)
	CheckTest(words []*models.Word, test_name, username string) ([]*models.CheckTestWord, error)
}

type TestService struct {
	repo repository.Repo
}

func (s *TestService) GetWordsForTest(name string) ([]*models.ElemTest, error) {
	// get all words -> get words for this test -> randomly mix -> return need test

	// get all words
	words, err := s.repo.GetWords()
	if err != nil {
		return nil, err
	}

	// get words for this test
	testwords, err := s.repo.GetTestByName(name)
	if err != nil {
		return nil, err
	}

	// generate test with random mix
	tests := []*models.ElemTest{}

	for i, item := range testwords.Words {
		test := models.ElemTest{}
		test.Words = make([]*models.Word, 4)
		root_i, _ := rand.Int(rand.Reader, big.NewInt(4))

		test.Words = tools.RandomWords(words, i)
		test.Words = tools.InsertByIndex(test.Words, int(root_i.Int64()), item)
		test.Right = int(root_i.Int64())

		tests = append(tests, &test)
	}
	return tests, nil
}

func (s *TestService) CheckTest(words []*models.Word, test_name, username string) ([]*models.CheckTestWord, error) {
	// get right words -> count percent, set time -> return right, wrong words, percent, time

	// get right words and generate map with this words
	answers, err := s.repo.GetWordsByNames(words)
	if err != nil {
		return nil, err
	}
	answersMap := map[string]string{}
	for _, item := range answers {
		answersMap[item.Word] = item.Translate
	}

	// generate result model and count right words
	test_answers := []*models.CheckTestWord{}

	percent_i := 0

	for _, item := range words {
		test_elm := models.CheckTestWord{
			Word:  item,
			Check: item.Translate == answersMap[item.Word],
			Right: answersMap[item.Word],
		}
		if test_elm.Check {
			percent_i++
		}
		test_answers = append(test_answers, &test_elm)
	}

	// count percent
	percent := math.Round((float64(percent_i) * 100 / float64(len(test_answers))))

	// set time
	time_res := time.Now().Format(time.ANSIC)

	// add result in test
	res := models.UserResult{Login: username, Percent: int(percent), Res: test_answers, Datatime: time_res}
	if err = s.repo.AddUserRes(res, test_name); err != nil {
		return nil, err
	}
	return test_answers, nil
}

func (s *TestService) AddTest(test models.Test) error {
	// set time
	test.Datatime = time.Now().Format(time.ANSIC)

	// add test
	return s.repo.AddTest(test)
}

func (s *TestService) GetTestByName(name string) (models.Test, error) {
	return s.repo.GetTestByName(name)
}

func (s *TestService) GetTests() ([]*models.Test, error) {
	return s.repo.GetTests()
}
