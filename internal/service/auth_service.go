package service

import (
	"errors"
	"time"

	"github.com/giicoo/maratWebSite/internal/repository"
	"github.com/giicoo/maratWebSite/internal/service/auth"
	"github.com/giicoo/maratWebSite/models"
	hashFunc "github.com/giicoo/maratWebSite/pkg/hash_password"
)

type AuthFuncs interface {
	SingIn(u models.User) (string, error)
	SingUp(u models.User) (models.User, error)
}
type AuthService struct {
	repo      repository.Repo
	hashTools hashFunc.HashTools
}

func (s *AuthService) SingIn(u models.User) (string, error) {
	// get user by login
	userInDB, err := s.repo.GetUserByLogin(u.Login)
	if err != nil {
		return "", err
	}
	// check password
	status := s.hashTools.CheckPasswordHash(u.Password, userInDB.Password)
	if status {
		return auth.NewJWT(u.Login)
	}
	return "", errors.New("Passwords is different")
}

func (s *AuthService) SingUp(u models.User) (models.User, error) {
	//hash password
	hash, err := s.hashTools.HashPassword(u.Password)
	if err != nil {
		return models.User{}, err
	}
	u.Password = hash

	// set data time
	u.Datatime = time.Now().Format(time.ANSIC)

	// add user
	return u, s.repo.AddUser(u)
}
