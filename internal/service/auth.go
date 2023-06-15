package service

import (
	"errors"

	"github.com/giicoo/maratWebSite/internal/service/auth"
	"github.com/giicoo/maratWebSite/models"
	hashFunc "github.com/giicoo/maratWebSite/pkg/hash_password"
)

func (s *Services) SingIn(u models.User) (string, error) {
	userInDB, err := s.repo.GetUser(u.Login)
	if err != nil {
		return "", err
	}
	status := hashFunc.CheckPasswordHash(u.Password, userInDB.Password)
	if status {
		return auth.NewJWT(u.Login)
	}
	return "", errors.New("Passwords is different")
}

func (s *Services) SingUp(u models.User) error {
	// validate |
	return s.repo.AddUser(u.Login, u.Password)
}
