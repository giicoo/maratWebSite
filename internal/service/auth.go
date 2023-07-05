package service

import (
	"errors"
	"time"

	"github.com/giicoo/maratWebSite/internal/repository"
	"github.com/giicoo/maratWebSite/internal/service/auth"
	"github.com/giicoo/maratWebSite/models"
	hashFunc "github.com/giicoo/maratWebSite/pkg/hash_password"
)

type AuthServices interface {
	SingIn(u models.UserDB) (string, error)
	SingUp(u models.UserDB) (models.UserDB, error)
}
type AuthService struct {
	repo      repository.Repo
	hashTools hashFunc.HashTools
}

func (s *AuthService) SingIn(u models.UserDB) (string, error) {
	userInDB, err := s.repo.GetUser(u.Login)
	if err != nil {
		return "", err
	}
	status := s.hashTools.CheckPasswordHash(u.Password, userInDB.Password)
	if status {
		return auth.NewJWT(u.Login)
	}
	return "", errors.New("Passwords is different")
}

func (s *AuthService) SingUp(u models.UserDB) (models.UserDB, error) {
	//hash password
	hash, err := s.hashTools.HashPassword(u.Password)
	if err != nil {
		return models.UserDB{}, err
	}
	u.Password = hash

	u.Datatime = time.Now().Format(time.ANSIC)

	return u, s.repo.AddUser(u)
}
