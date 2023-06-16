package repository

import "github.com/giicoo/maratWebSite/models"

type Repo interface {
	AddUser(login string, hash_password string) error
	GetUser(login string) (models.User, error)
	AddWord(word string, translate string) error
	GetWords() ([]models.Word, error)
}
