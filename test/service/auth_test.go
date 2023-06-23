package http_v1_test

import (
	"errors"
	"testing"

	mock_repository "github.com/giicoo/maratWebSite/internal/repository/mocks"
	"github.com/giicoo/maratWebSite/internal/service"
	"github.com/giicoo/maratWebSite/models"
	mock_hashFunc "github.com/giicoo/maratWebSite/pkg/hash_password/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSingUp(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, user models.UserDB)
	type mockBehaviorHash func(r *mock_hashFunc.MockHashTools, password string)

	tests := []struct {
		name             string
		inputUser        models.UserDB
		inputUserDB      models.UserDB
		mockBehavior     mockBehavior
		mockBehaviorHash mockBehaviorHash
		expectedError    error
		expectedUser     models.UserDB
	}{
		{
			name:        "OK",
			inputUser:   models.UserDB{Login: "test", Password: "test_p"},
			inputUserDB: models.UserDB{Login: "test", Password: "hash"},
			mockBehavior: func(r *mock_repository.MockRepo, user models.UserDB) {
				r.EXPECT().AddUser(user).Return(nil)
			},
			mockBehaviorHash: func(r *mock_hashFunc.MockHashTools, password string) {
				r.EXPECT().HashPassword(password).Return("hash", nil)
			},
			expectedError: nil,
			expectedUser:  models.UserDB{Login: "test", Password: "hash"},
		},
		{
			name:        "DB Error",
			inputUser:   models.UserDB{Login: "test", Password: "test_p"},
			inputUserDB: models.UserDB{Login: "test", Password: "hash"},
			mockBehavior: func(r *mock_repository.MockRepo, user models.UserDB) {
				r.EXPECT().AddUser(user).Return(errors.New("DB Error"))
			},
			mockBehaviorHash: func(r *mock_hashFunc.MockHashTools, password string) {
				r.EXPECT().HashPassword(password).Return("hash", nil)
			},
			expectedError: errors.New("DB Error"),
			expectedUser:  models.UserDB{Login: "test", Password: "hash"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init mock repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo, test.inputUserDB)

			hash := mock_hashFunc.NewMockHashTools(c)
			test.mockBehaviorHash(hash, test.inputUser.Password)

			// init services
			services := service.NewServices(repo, hash)

			user, err := services.AuthServices.SingUp(test.inputUser)

			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedUser, user)
		})
	}
}
