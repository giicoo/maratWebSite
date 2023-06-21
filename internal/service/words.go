package service

import (
	"math/rand"
	"time"

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
		test.Words = randomWords(words, i)
		test.Words = insert(test.Words, root_i, item)
		test.Right = root_i

		tests = append(tests, &test)
	}
	return tests, nil
}

func randomWords(w []*models.WordDB, i_root int) []*models.WordDB {
	words := make([]*models.WordDB, len(w))
	copy(words, w)

	words = words[:]
	words[i_root], words[len(words)-1] = words[len(words)-1], words[i_root]

	words = words[:len(words)-1]
	rand.Seed(time.Now().Unix())
	i1 := rand.Intn(len(words))
	w1 := words[i1]

	words[i1], words[len(words)-1] = words[len(words)-1], words[i1]

	words = words[:len(words)-1]
	i2 := rand.Intn(len(words))
	w2 := words[i2]

	words[i2], words[len(words)-1] = words[len(words)-1], words[i2]

	words = words[:len(words)-1]
	i3 := rand.Intn(len(words))
	w3 := words[i3]

	words[i3], words[len(words)-1] = words[len(words)-1], words[i3]

	return []*models.WordDB{w1, w2, w3}
}

func insert(a []*models.WordDB, index int, value *models.WordDB) []*models.WordDB {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
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
