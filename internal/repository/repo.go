package repository

import "github.com/giicoo/maratWebSite/models"

type Repo interface {
	AddUser(user models.UserDB) error
	GetUser(login string) (models.UserDB, error)
	AddWord(word models.WordDB) error
	GetWords() ([]*models.WordDB, error)
}
