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

func TestGetWords(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo)

	tests := []struct {
		name           string
		mockBehavior   mockBehavior
		expectedError  error
		expectedAnswer []*models.WordDB
	}{
		{
			name: "OK",
			mockBehavior: func(r *mock_repository.MockRepo) {
				r.EXPECT().GetWords().Return([]*models.WordDB{{Word: "test", Translate: "test_t"}}, nil)
			},
			expectedError:  nil,
			expectedAnswer: []*models.WordDB{{Word: "test", Translate: "test_t"}},
		},
		{
			name: "Empty",
			mockBehavior: func(r *mock_repository.MockRepo) {
				r.EXPECT().GetWords().Return(nil, nil)
			},
			expectedError:  nil,
			expectedAnswer: nil,
		},
		{
			name: "Empty",
			mockBehavior: func(r *mock_repository.MockRepo) {
				r.EXPECT().GetWords().Return(nil, errors.New("Test Error"))
			},
			expectedError:  errors.New("Test Error"),
			expectedAnswer: nil,
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

			// check services
			words, err := services.WordsServices.GetWord()

			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedAnswer, words)
		})
	}
}

func TestAddWord(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, word models.WordDB)

	tests := []struct {
		name          string
		inputWord     models.WordDB
		mockBehavior  mockBehavior
		expectedError error
	}{
		{
			name:      "OK",
			inputWord: models.WordDB{Word: "test", Translate: "test_t"},
			mockBehavior: func(r *mock_repository.MockRepo, word models.WordDB) {
				r.EXPECT().AddWord(word).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:      "Service Error",
			inputWord: models.WordDB{Word: "test", Translate: "test_t"},
			mockBehavior: func(r *mock_repository.MockRepo, word models.WordDB) {
				r.EXPECT().AddWord(word).Return(errors.New("Test Error"))
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
			test.mockBehavior(repo, test.inputWord)

			hash := hashFunc.NewHashTools()
			// init services
			services := service.NewServices(repo, hash)

			// check service
			err := services.WordsServices.AddWord(test.inputWord)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
