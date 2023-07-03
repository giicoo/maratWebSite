package repository

import "github.com/giicoo/maratWebSite/models"

//go:generate $GOPATH/bin/mockgen -source=repo.go -destination=mocks/mock.go
type Repo interface {
	AddUser(user models.UserDB) error
	GetUser(login string) (models.UserDB, error)
	AddWord(word models.WordDB) error
	GetWords() ([]*models.WordDB, error)
	GetWordsByNames(words []*models.WordDB) ([]*models.WordDB, error)
	AddTest(test models.Test) error
	GetTestByName(name string) (models.Test, error)
}
