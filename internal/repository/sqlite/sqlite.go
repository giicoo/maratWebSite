package sqlite

import (
	"database/sql"
	"os"
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
	if err != nil {
		return err
	}

	stmt, err = os.ReadFile("internal/repository/sqlite/sql/create-table_1.sql")
	if err != nil {
		return err
	}
	_, err = s.db.Exec(string(stmt))
	if err != nil {
		return err
	}

	return nil
}
