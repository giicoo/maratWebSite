package http_v1_test

import (
	"errors"
	"testing"

	mock_repository "github.com/giicoo/maratWebSite/internal/repository/mocks"
	"github.com/giicoo/maratWebSite/internal/service"
	"github.com/giicoo/maratWebSite/models"
	hashFunc "github.com/giicoo/maratWebSite/pkg/hash_password"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetWordsForTest(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo)

	tests := []struct {
		name          string
		mockBehavior  mockBehavior
		expectedError error
		expectedEmpty bool
	}{
		{
			name: "OK",
			mockBehavior: func(r *mock_repository.MockRepo) {
				r.EXPECT().GetWords().Return([]*models.WordDB{{Word: "test", Translate: "test_t"}, {Word: "test1", Translate: "test_t1"}, {Word: "test2", Translate: "test_t2"}, {Word: "test3", Translate: "test_t3"}}, nil)
			},
			expectedError: nil,
			expectedEmpty: false,
		},
		{
			name: "Service Error",
			mockBehavior: func(r *mock_repository.MockRepo) {
				r.EXPECT().GetWords().Return(nil, errors.New("Test Error"))
			},
			expectedError: errors.New("Test Error"),
			expectedEmpty: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init mock repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo)

			hash := hashFunc.NewHashTools()
			// init services
			services := service.NewServices(repo, hash)

			words, err := services.WordsServices.GetWordsForTest()

			assert.Equal(t, test.expectedError, err)

			if test.expectedEmpty {
				assert.Empty(t, words)
			} else {
				assert.NotEmpty(t, words)
			}
		})
	}
}

func TestCheckTest(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, words []*models.WordDB)

	tests := []struct {
		name           string
		inputWords     []*models.WordDB
		mockBehavior   mockBehavior
		expectedError  error
		expectedAnswer []*models.TestWord
	}{
		{
			name:       "OK",
			inputWords: []*models.WordDB{{Word: "test1", Translate: "test_t1"}},
			mockBehavior: func(r *mock_repository.MockRepo, words []*models.WordDB) {
				r.EXPECT().GetWordsByNames(words).Return([]*models.WordDB{{Word: "test1", Translate: "test_t1"}}, nil)
			},
			expectedError:  nil,
			expectedAnswer: []*models.TestWord{{Word: &models.WordDB{Word: "test1", Translate: "test_t1"}, Check: true, Right: "test_t1"}},
		},
		{
			name:       "OK_with_wrong_answer",
			inputWords: []*models.WordDB{{Word: "test1", Translate: "test_t"}},
			mockBehavior: func(r *mock_repository.MockRepo, words []*models.WordDB) {
				r.EXPECT().GetWordsByNames(words).Return([]*models.WordDB{{Word: "test1", Translate: "test_t1"}}, nil)
			},
			expectedError:  nil,
			expectedAnswer: []*models.TestWord{{Word: &models.WordDB{Word: "test1", Translate: "test_t"}, Check: false, Right: "test_t1"}},
		},
		{
			name:       "Service Error",
			inputWords: []*models.WordDB{{Word: "test1", Translate: "test_t1"}, {Word: "test2", Translate: "test_t2"}},
			mockBehavior: func(r *mock_repository.MockRepo, words []*models.WordDB) {
				r.EXPECT().GetWordsByNames(words).Return(nil, errors.New("Test Error"))
			},
			expectedError: errors.New("Test Error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init mock repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo, test.inputWords)

			hash := hashFunc.NewHashTools()
			// init services
			services := service.NewServices(repo, hash)

			words, err := services.WordsServices.CheckTest(test.inputWords)
			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedAnswer, words)
		})
	}
}
