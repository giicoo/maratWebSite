package http_v1_test

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	http_v1 "github.com/giicoo/maratWebSite/internal/delivery/http/v1"
	mock_repository "github.com/giicoo/maratWebSite/internal/repository/mocks"
	"github.com/giicoo/maratWebSite/internal/service"
	"github.com/giicoo/maratWebSite/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetWords(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo)

	tests := []struct {
		name           string
		mockBehavior   mockBehavior
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "OK",
			mockBehavior: func(r *mock_repository.MockRepo) {
				r.EXPECT().GetWords().Return([]*models.WordDB{{Word: "test", Translate: "test_t"}}, nil)
			},
			expectedStatus: 200,
			expectedBody:   `[{"word":"test","translate":"test_t"}]`,
		},
		{
			name: "Empty",
			mockBehavior: func(r *mock_repository.MockRepo) {
				r.EXPECT().GetWords().Return(nil, nil)
			},
			expectedStatus: 200,
			expectedBody:   "null",
		},
		{
			name: "Empty",
			mockBehavior: func(r *mock_repository.MockRepo) {
				r.EXPECT().GetWords().Return(nil, errors.New(""))
			},
			expectedStatus: 500,
			expectedBody:   "Service Error\nnull",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init mock repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo)

			// init services
			services := service.NewServices(repo)

			// init handler and router
			h := http_v1.NewHandler(services)
			r := h.InitHandlers()

			// serve request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/get-words", bytes.NewBufferString(""))
			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatus, w.Code)
			assert.Equal(t, test.expectedBody, w.Body.String())
		})
	}
}

func TestAddWord(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, word models.WordDB)

	tests := []struct {
		name           string
		inputBody      string
		inputWord      models.WordDB
		mockBehavior   mockBehavior
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "OK",
			inputBody: `{"word":"test","translate":"test_t"}`,
			inputWord: models.WordDB{Word: "test", Translate: "test_t"},
			mockBehavior: func(r *mock_repository.MockRepo, word models.WordDB) {
				r.EXPECT().AddWord(word).Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   "Successful Add",
		},
		{
			name:           "Invalid JSON",
			inputBody:      ``,
			inputWord:      models.WordDB{Word: "test", Translate: "test_t"},
			mockBehavior:   func(r *mock_repository.MockRepo, word models.WordDB) {},
			expectedStatus: 400,
			expectedBody:   "Invalid JSON\n",
		},
		{
			name:      "Service Error",
			inputBody: `{"word":"test","translate":"test_t"}`,
			inputWord: models.WordDB{Word: "test", Translate: "test_t"},
			mockBehavior: func(r *mock_repository.MockRepo, word models.WordDB) {
				r.EXPECT().AddWord(word).Return(errors.New(""))
			},
			expectedStatus: 500,
			expectedBody:   "Service Error\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init mock repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo, test.inputWord)

			// init services
			services := service.NewServices(repo)

			// init handler and router
			h := http_v1.NewHandler(services)
			r := h.InitHandlers()

			// serve request

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/add-word", bytes.NewBufferString(test.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatus, w.Code)
			assert.Equal(t, test.expectedBody, w.Body.String())
		})
	}
}
