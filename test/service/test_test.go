package http_v1_test

import (
	"errors"
	"testing"
	"time"

	"github.com/giicoo/maratWebSite/configs"
	mock_repository "github.com/giicoo/maratWebSite/internal/repository/mocks"
	"github.com/giicoo/maratWebSite/internal/service"
	"github.com/giicoo/maratWebSite/models"
	hashFunc "github.com/giicoo/maratWebSite/pkg/hash_password"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetTests(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo)

	tests := []struct {
		name           string
		mockBehavior   mockBehavior
		expectedError  error
		expectedAnswer []*models.Test
	}{
		{
			name: "OK",
			mockBehavior: func(r *mock_repository.MockRepo) {
				r.EXPECT().GetTests().Return([]*models.Test{{
					Name:         "test",
					Words:        []*models.Word{},
					UsersResults: []*models.UserResult{},
					Datatime:     "test"}}, nil)
			},
			expectedError: nil,
			expectedAnswer: []*models.Test{{
				Name:         "test",
				Words:        []*models.Word{},
				UsersResults: []*models.UserResult{},
				Datatime:     "test"}},
		},
		{
			name: "Error",
			mockBehavior: func(r *mock_repository.MockRepo) {
				r.EXPECT().GetTests().Return(nil, errors.New("test err"))
			},
			expectedError:  errors.New("test err"),
			expectedAnswer: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo)

			// init services
			hash := hashFunc.NewHashTools()
			cfg := &configs.Config{ADMIN_LOGIN: "admin", TIME_COOKIE: 3600}
			services := service.NewServices(repo, hash, cfg)

			tests, err := services.TestServices.GetTests()

			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedAnswer, tests)
		})
	}
}

func TestAddTest(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, test models.Test)

	tests := []struct {
		name          string
		mockBehavior  mockBehavior
		inputTest     models.Test
		expectedError error
	}{
		{
			name: "OK",
			inputTest: models.Test{
				Name:         "Test",
				Words:        []*models.Word{{Word: "test", Translate: "test_t"}},
				UsersResults: []*models.UserResult{},
				Datatime:     time.Now().Format(time.ANSIC),
			},
			mockBehavior: func(r *mock_repository.MockRepo, test models.Test) {
				r.EXPECT().AddTest(test).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error",
			inputTest: models.Test{
				Name:         "Test",
				Words:        []*models.Word{{Word: "test", Translate: "test_t"}},
				UsersResults: []*models.UserResult{},
				Datatime:     time.Now().Format(time.ANSIC),
			},
			mockBehavior: func(r *mock_repository.MockRepo, test models.Test) {
				r.EXPECT().AddTest(test).Return(errors.New("test err"))
			},
			expectedError: errors.New("test err"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo, test.inputTest)

			// init service
			hash := hashFunc.NewHashTools()
			cfg := &configs.Config{ADMIN_LOGIN: "admin", TIME_COOKIE: 3600}
			services := service.NewServices(repo, hash, cfg)

			err := services.TestServices.AddTest(test.inputTest)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
func TestDeleteTest(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, test models.Test)

	tests := []struct {
		name          string
		mockBehavior  mockBehavior
		inputTest     []models.Test
		expectedError error
	}{
		{
			name: "OK",
			inputTest: []models.Test{{
				Name:         "Test",
				Words:        []*models.Word{{Word: "test", Translate: "test_t"}},
				UsersResults: []*models.UserResult{},
				Datatime:     time.Now().Format(time.ANSIC),
			}},
			mockBehavior: func(r *mock_repository.MockRepo, test models.Test) {
				r.EXPECT().DeleteTest(test).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error",
			inputTest: []models.Test{{
				Name:         "Test",
				Words:        []*models.Word{{Word: "test", Translate: "test_t"}},
				UsersResults: []*models.UserResult{},
				Datatime:     time.Now().Format(time.ANSIC),
			}},
			mockBehavior: func(r *mock_repository.MockRepo, test models.Test) {
				r.EXPECT().DeleteTest(test).Return(errors.New("test err"))
			},
			expectedError: errors.New("test err"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo, test.inputTest[0])

			// init service
			hash := hashFunc.NewHashTools()
			cfg := &configs.Config{ADMIN_LOGIN: "admin", TIME_COOKIE: 3600}
			services := service.NewServices(repo, hash, cfg)

			err := services.TestServices.DeleteTest(test.inputTest)

			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestGetTestByName(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, name string)

	tests := []struct {
		name          string
		nameForTest   string
		mockBehavior  mockBehavior
		expectedError error
		expectedTest  models.Test
	}{
		{
			name:        "OK",
			nameForTest: "test",
			mockBehavior: func(r *mock_repository.MockRepo, name string) {
				r.EXPECT().GetTestByName(name).Return(models.Test{Name: "test"}, nil)
			},
			expectedError: nil,
			expectedTest:  models.Test{Name: "test"},
		},
		{
			name:        "Error",
			nameForTest: "test",
			mockBehavior: func(r *mock_repository.MockRepo, name string) {
				r.EXPECT().GetTestByName(name).Return(models.Test{}, errors.New("test err"))
			},
			expectedError: errors.New("test err"),
			expectedTest:  models.Test{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo, test.nameForTest)

			// init service
			hash := hashFunc.NewHashTools()
			cfg := &configs.Config{ADMIN_LOGIN: "admin", TIME_COOKIE: 3600}
			services := service.NewServices(repo, hash, cfg)

			testFromDB, err := services.TestServices.GetTestByName(test.nameForTest)

			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedTest, testFromDB)
		})
	}
}

func TestGetWordsForTest(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, name string)

	tests := []struct {
		name          string
		nameTest      string
		mockBehavior  mockBehavior
		expectedError error
	}{
		{
			name:     "OK",
			nameTest: "Test",
			mockBehavior: func(r *mock_repository.MockRepo, name string) {
				r.EXPECT().GetWords().Return([]*models.Word{{Word: "test", Translate: "test_t"}}, nil)
				r.EXPECT().GetTestByName(name).Return(models.Test{Name: "Test", Words: []*models.Word{}}, nil)
			},
			expectedError: nil,
		},
		{
			name:     "Error GetWords",
			nameTest: "Test",
			mockBehavior: func(r *mock_repository.MockRepo, name string) {
				r.EXPECT().GetWords().Return([]*models.Word{{Word: "test", Translate: "test_t"}}, errors.New("test err"))
			},
			expectedError: errors.New("test err"),
		},
		{
			name:     "Error GetTestByName",
			nameTest: "Test",
			mockBehavior: func(r *mock_repository.MockRepo, name string) {
				r.EXPECT().GetWords().Return([]*models.Word{{Word: "test", Translate: "test_t"}}, nil)
				r.EXPECT().GetTestByName(name).Return(models.Test{}, errors.New("test err"))
			},
			expectedError: errors.New("test err"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init mock repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo, test.nameTest)

			// init services
			hash := hashFunc.NewHashTools()
			cfg := &configs.Config{ADMIN_LOGIN: "admin", TIME_COOKIE: 3600}
			services := service.NewServices(repo, hash, cfg)

			_, err := services.TestServices.GetWordsForTest(test.nameTest)

			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestCheckTest(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockRepo, words []*models.Word, res models.UserResult, username string)

	tests := []struct {
		name           string
		nameTest       string
		username       string
		userRes        models.UserResult
		inputWords     []*models.Word
		mockBehavior   mockBehavior
		expectedError  error
		expectedAnswer []*models.CheckTestWord
	}{
		{
			name:     "OK",
			nameTest: "Test",
			username: "test",
			userRes: models.UserResult{
				Login:   "test",
				Percent: 100,
				Res: []*models.CheckTestWord{{
					Word:  &models.Word{Word: "test1", Translate: "test_t1"},
					Check: true, Right: "test_t1"}},
				Datatime: time.Now().Format(time.ANSIC)},
			inputWords: []*models.Word{{Word: "test1", Translate: "test_t1"}},
			mockBehavior: func(r *mock_repository.MockRepo, words []*models.Word, res models.UserResult, test_name string) {
				r.EXPECT().GetWordsByNames(words).Return([]*models.Word{{Word: "test1", Translate: "test_t1"}}, nil)
				r.EXPECT().AddUserRes(res, test_name).Return(nil)
			},
			expectedError:  nil,
			expectedAnswer: []*models.CheckTestWord{{Word: &models.Word{Word: "test1", Translate: "test_t1"}, Check: true, Right: "test_t1"}},
		},
		{
			name:     "Wrong Answer",
			nameTest: "Test",
			username: "test",
			userRes: models.UserResult{
				Login:   "test",
				Percent: 0,
				Res: []*models.CheckTestWord{{
					Word:  &models.Word{Word: "test1", Translate: "test_t1"},
					Check: false, Right: "test_t"}},
				Datatime: time.Now().Format(time.ANSIC)},
			inputWords: []*models.Word{{Word: "test1", Translate: "test_t1"}},
			mockBehavior: func(r *mock_repository.MockRepo, words []*models.Word, res models.UserResult, test_name string) {
				r.EXPECT().GetWordsByNames(words).Return([]*models.Word{{Word: "test1", Translate: "test_t"}}, nil)
				r.EXPECT().AddUserRes(res, test_name).Return(nil)
			},
			expectedError:  nil,
			expectedAnswer: []*models.CheckTestWord{{Word: &models.Word{Word: "test1", Translate: "test_t1"}, Check: false, Right: "test_t"}},
		},
		{
			name:       "Error_1",
			nameTest:   "Test",
			username:   "test",
			userRes:    models.UserResult{},
			inputWords: []*models.Word{},
			mockBehavior: func(r *mock_repository.MockRepo, words []*models.Word, res models.UserResult, test_name string) {
				r.EXPECT().GetWordsByNames(words).Return(nil, errors.New("test err"))
			},
			expectedError:  errors.New("test err"),
			expectedAnswer: nil,
		},
		{
			name:     "Error_2",
			nameTest: "Test",
			username: "test",
			userRes: models.UserResult{
				Login:   "test",
				Percent: 100,
				Res: []*models.CheckTestWord{{
					Word:  &models.Word{Word: "test1", Translate: "test_t1"},
					Check: true, Right: "test_t1"}},
				Datatime: time.Now().Format(time.ANSIC)},
			inputWords: []*models.Word{{Word: "test1", Translate: "test_t1"}},
			mockBehavior: func(r *mock_repository.MockRepo, words []*models.Word, res models.UserResult, test_name string) {
				r.EXPECT().GetWordsByNames(words).Return([]*models.Word{{Word: "test1", Translate: "test_t1"}}, nil)
				r.EXPECT().AddUserRes(res, test_name).Return(errors.New("test err"))
			},
			expectedError:  errors.New("test err"),
			expectedAnswer: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			// init mock repo
			repo := mock_repository.NewMockRepo(c)
			test.mockBehavior(repo, test.inputWords, test.userRes, test.nameTest)

			hash := hashFunc.NewHashTools()
			cfg := &configs.Config{ADMIN_LOGIN: "admin", TIME_COOKIE: 3600}
			// init services
			services := service.NewServices(repo, hash, cfg)

			words, err := services.TestServices.CheckTest(test.inputWords, test.nameTest, test.username)
			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedAnswer, words)
		})
	}
}
