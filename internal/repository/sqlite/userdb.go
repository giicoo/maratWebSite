package sqlite

import "github.com/giicoo/maratWebSite/models"

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
