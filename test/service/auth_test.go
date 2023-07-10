package http_v1_test

import (
	"errors"
	"testing"
	"time"

	mock_repository "github.com/giicoo/maratWebSite/internal/repository/mocks"
	"github.com/giicoo/maratWebSite/internal/service"
	"github.com/giicoo/maratWebSite/models"
	mock_hashFunc "github.com/giicoo/maratWebSite/pkg/hash_password/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSingUp(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, user models.User)
	type mockBehaviorHash func(r *mock_hashFunc.MockHashTools, password string)

	tests := []struct {
		name             string
		inputUser        models.User
		inputUserDB      models.User
		mockBehavior     mockBehavior
		mockBehaviorHash mockBehaviorHash
		expectedError    error
		expectedUser     models.User
	}{
		{
			name:        "OK",
			inputUser:   models.User{Login: "test", Password: "test_p", Datatime: time.Now().Format(time.ANSIC)},
			inputUserDB: models.User{Login: "test", Password: "hash", Datatime: time.Now().Format(time.ANSIC)},
			mockBehavior: func(r *mock_repository.MockRepo, user models.User) {
				r.EXPECT().AddUser(user).Return(nil)
			},
			mockBehaviorHash: func(r *mock_hashFunc.MockHashTools, password string) {
				r.EXPECT().HashPassword(password).Return("hash", nil)
			},
			expectedError: nil,
			expectedUser:  models.User{Login: "test", Password: "hash", Datatime: time.Now().Format(time.ANSIC)},
		},
		{
			name:        "DB Error",
			inputUser:   models.User{Login: "test", Password: "test_p", Datatime: time.Now().Format(time.ANSIC)},
			inputUserDB: models.User{Login: "test", Password: "hash", Datatime: time.Now().Format(time.ANSIC)},
			mockBehavior: func(r *mock_repository.MockRepo, user models.User) {
				r.EXPECT().AddUser(user).Return(errors.New("DB Error"))
			},
			mockBehaviorHash: func(r *mock_hashFunc.MockHashTools, password string) {
				r.EXPECT().HashPassword(password).Return("hash", nil)
			},
			expectedError: errors.New("DB Error"),
			expectedUser:  models.User{Login: "test", Password: "hash", Datatime: time.Now().Format(time.ANSIC)},
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

func TestSingIn(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, login string)
	type mockBehaviorHash func(r *mock_hashFunc.MockHashTools, password string, hash string)

	tests := []struct {
		name             string
		inputUser        models.User
		inputUserDB      models.User
		mockBehavior     mockBehavior
		mockBehaviorHash mockBehaviorHash
		expectedError    error
		expectedAnswer   string
	}{
		{
			name:        "OK",
			inputUser:   models.User{Login: "test", Password: "testp", Datatime: time.Now().Format(time.ANSIC)},
			inputUserDB: models.User{Login: "test", Password: "testp", Datatime: time.Now().Format(time.ANSIC)},
			mockBehavior: func(r *mock_repository.MockRepo, login string) {
				r.EXPECT().GetUserByLogin(login).Return(models.User{Login: "test", Password: "testp", Datatime: time.Now().Format(time.ANSIC)}, nil)
			},
			mockBehaviorHash: func(r *mock_hashFunc.MockHashTools, password, hash string) {
				r.EXPECT().CheckPasswordHash(password, hash).Return(true)
			},
			expectedError: nil,
		},
		{
			name:        "Invalid Password",
			inputUser:   models.User{Login: "test", Password: "test", Datatime: time.Now().Format(time.ANSIC)},
			inputUserDB: models.User{Login: "test", Password: "testp", Datatime: time.Now().Format(time.ANSIC)},
			mockBehavior: func(r *mock_repository.MockRepo, login string) {
				r.EXPECT().GetUserByLogin(login).Return(models.User{Login: "test", Password: "testp", Datatime: time.Now().Format(time.ANSIC)}, nil)
			},
			mockBehaviorHash: func(r *mock_hashFunc.MockHashTools, password, hash string) {
				r.EXPECT().CheckPasswordHash(password, hash).Return(false)
			},
			expectedError: errors.New("Passwords is different"),
		},
		{
			name:      "DB Error",
			inputUser: models.User{Login: "test", Password: "test", Datatime: time.Now().Format(time.ANSIC)},
			mockBehavior: func(r *mock_repository.MockRepo, login string) {
				r.EXPECT().GetUserByLogin(login).Return(models.User{Login: "", Password: ""}, errors.New("test error"))
			},
			mockBehaviorHash: func(r *mock_hashFunc.MockHashTools, password, hash string) {},
			expectedError:    errors.New("test error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo, test.inputUser.Login)

			hash := mock_hashFunc.NewMockHashTools(c)
			test.mockBehaviorHash(hash, test.inputUser.Password, test.inputUserDB.Password)

			services := service.NewServices(repo, hash)
			_, err := services.AuthServices.SingIn(test.inputUser)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
