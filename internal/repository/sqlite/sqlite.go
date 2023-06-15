package sqlite

import (
	"database/sql"
	"os"

	"github.com/giicoo/maratWebSite/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) InitDB() error {
	if err := s.db.Ping(); err != nil {
		return err
	}

	stmt, err := os.ReadFile("internal/repository/sqlite/sql/create-table.sql")
	if err != nil {
		return err
	}
	_, err = s.db.Exec(string(stmt))
	return err
}

func (s *Store) AddUser(login string, hash_password string) error {
	stmt := "INSERT INTO users(login, hash_password) VALUES(@param1, @param2)"
	_, err := s.db.Exec(stmt, login, hash_password)
	return err
}

func (s *Store) GetUser(login string) (models.User, error) {
	stmt := "SELECT login, hash_password FROM users WHERE login=@param1"
	row := s.db.QueryRow(stmt, login)
	user := models.User{}

	err := row.Scan(&user.Login, &user.Password)
	return user, err
}
