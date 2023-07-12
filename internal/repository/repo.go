package repository

import "github.com/giicoo/maratWebSite/models"

//go:generate $GOPATH/bin/mockgen -source=repo.go -destination=mocks/mock.go
type Repo interface {
	// user
	AddUser(user models.User) error
	GetUserByLogin(login string) (models.User, error)

	//words
	AddWord(word models.Word) error
	DeleteWord(w models.Word) error
	GetWords() ([]*models.Word, error)
	GetWordsByNames(words []*models.Word) ([]*models.Word, error)

	// test
	AddTest(test models.Test) error
	GetTestByName(name string) (models.Test, error)
	AddUserRes(res models.UserResult, test_name string) error
	GetTests() ([]*models.Test, error)
	DeleteTest(w models.Test) error
}
