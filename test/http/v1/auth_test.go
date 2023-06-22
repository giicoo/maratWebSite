package http_v1_test

import (
	"bytes"
	"net/http/httptest"
	"testing"

	http_v1 "github.com/giicoo/maratWebSite/internal/delivery/http/v1"
	mock_repository "github.com/giicoo/maratWebSite/internal/repository/mocks"
	"github.com/giicoo/maratWebSite/internal/service"
	"github.com/giicoo/maratWebSite/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSingUp(t *testing.T) {
	//TODO: IS NOT WORK, BECAUSE SERVICE HASH PASSWORD, WATCH VIDEO!!!
	type mockBehavior func(r *mock_repository.MockRepo, user models.UserDB)

	tests := []struct {
		name           string
		inputBody      string
		inputUser      models.UserDB
		mockBehavior   mockBehavior
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "OK",
			inputBody: `{"login":"test","password":"test_p"}`,
			inputUser: models.UserDB{Login: "test", Password: "test_p"},
			mockBehavior: func(r *mock_repository.MockRepo, user models.UserDB) {
				r.EXPECT().AddUser(user).Return(nil)
			},
			expectedStatus: 200,
			expectedBody:   "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init mock repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo, test.inputUser)

			// init services
			services := service.NewServices(repo)

			// init handler and router
			h := http_v1.NewHandler(services)
			r := h.InitHandlers()

			// serve request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/singup", bytes.NewBufferString(test.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatus, w.Code)
			assert.Equal(t, test.expectedBody, w.Body.String())
		})
	}
}
