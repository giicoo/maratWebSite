package repository

import "github.com/giicoo/maratWebSite/models"

type Repo interface {
	AddUser(login string, hash_password string) error
	GetUser(login string) (models.User, error)
}
