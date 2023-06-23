package hashFunc

import "golang.org/x/crypto/bcrypt"

//go:generate $GOPATH/bin/mockgen -source=hashing.go -destination=mocks/mock.go
type HashTools interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
type Hash struct{}

func NewHashTools() *Hash {
	return &Hash{}
}
func (h *Hash) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h *Hash) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
