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

func TestTestPage(t *testing.T) {
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
				r.EXPECT().GetWords().Return([]*models.WordDB{{Word: "test", Translate: "test_t"}, {Word: "test1", Translate: "test_t1"}, {Word: "test2", Translate: "test_t2"}, {Word: "test3", Translate: "test_t3"}}, nil)
			},
			expectedStatus: 200,
		},
		{
			name: "Service Error",
			mockBehavior: func(r *mock_repository.MockRepo) {
				r.EXPECT().GetWords().Return(nil, errors.New(""))
			},
			expectedStatus: 500,
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
			req := httptest.NewRequest("POST", "/get-words-for-test", bytes.NewBufferString(""))
			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatus, w.Code)
		})
	}
}

func TestCheckTest(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, words []*models.WordDB)

	tests := []struct {
		name           string
		inputBody      string
		inputWords     []*models.WordDB
		mockBehavior   mockBehavior
		expectedStatus int
		expectedBody   string
	}{
		{
			name:       "OK",
			inputBody:  `[{"word":"test1","translate":"test_t1"},{"word":"test2","translate":"test_t2"}]`,
			inputWords: []*models.WordDB{{Word: "test1", Translate: "test_t1"}, {Word: "test2", Translate: "test_t2"}},
			mockBehavior: func(r *mock_repository.MockRepo, words []*models.WordDB) {
				r.EXPECT().GetWordsByNames(words).Return([]*models.WordDB{{Word: "test1", Translate: "test_t1"}, {Word: "test2", Translate: "test_t2"}}, nil)
			},
			expectedStatus: 200,
			expectedBody:   "[{\"word\":{\"word\":\"test1\",\"translate\":\"test_t1\"},\"check\":true,\"right\":\"test_t1\"},{\"word\":{\"word\":\"test2\",\"translate\":\"test_t2\"},\"check\":true,\"right\":\"test_t2\"}]",
		},
		{
			name:       "OK_with_wrong_answer",
			inputBody:  `[{"word":"test1","translate":"test_t"},{"word":"test2","translate":"test_t2"}]`,
			inputWords: []*models.WordDB{{Word: "test1", Translate: "test_t"}, {Word: "test2", Translate: "test_t2"}},
			mockBehavior: func(r *mock_repository.MockRepo, words []*models.WordDB) {
				r.EXPECT().GetWordsByNames(words).Return([]*models.WordDB{{Word: "test1", Translate: "test_t1"}, {Word: "test2", Translate: "test_t2"}}, nil)
			},
			expectedStatus: 200,
			expectedBody:   "[{\"word\":{\"word\":\"test1\",\"translate\":\"test_t\"},\"check\":false,\"right\":\"test_t1\"},{\"word\":{\"word\":\"test2\",\"translate\":\"test_t2\"},\"check\":true,\"right\":\"test_t2\"}]",
		},
		{
			name:       "Invalid JSON",
			inputBody:  ``,
			inputWords: []*models.WordDB{{Word: "test1", Translate: "test_t"}, {Word: "test2", Translate: "test_t2"}},
			mockBehavior: func(r *mock_repository.MockRepo, words []*models.WordDB) {

			},
			expectedStatus: 400,
			expectedBody:   "Invalid JSON\n",
		},
		{
			name:       "Service Error",
			inputBody:  `[{"word":"test1","translate":"test_t1"},{"word":"test2","translate":"test_t2"}]`,
			inputWords: []*models.WordDB{{Word: "test1", Translate: "test_t1"}, {Word: "test2", Translate: "test_t2"}},
			mockBehavior: func(r *mock_repository.MockRepo, words []*models.WordDB) {
				r.EXPECT().GetWordsByNames(words).Return(nil, errors.New(""))
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
			test.mockBehavior(repo, test.inputWords)

			// init services
			services := service.NewServices(repo)

			// init handler and router
			h := http_v1.NewHandler(services)
			r := h.InitHandlers()

			// serve request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/check-test", bytes.NewBufferString(test.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatus, w.Code)
			assert.Equal(t, test.expectedBody, w.Body.String())
		})
	}
}
