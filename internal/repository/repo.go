package repository

import "github.com/giicoo/maratWebSite/models"

//go:generate $GOPATH/bin/mockgen -source=repo.go -destination=mocks/mock.go
type Repo interface {
	// user
	AddUser(user models.UserDB) error
	GetUser(login string) (models.UserDB, error)

	//words
	AddWord(word models.WordDB) error
	GetWords() ([]*models.WordDB, error)
	GetWordsByNames(words []*models.WordDB) ([]*models.WordDB, error)

	// test
	AddTest(test models.Test) error
	GetTestByName(name string) (models.Test, error)
	AddUserRes(res models.UserResult, test_name string) error
}
