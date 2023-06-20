package sqlite

import "github.com/giicoo/maratWebSite/models"

func (s *Store) AddWord(word, translate string) error {
	stmt := "INSERT INTO words(word, translate) VALUES(@param1, @param2)"
	_, err := s.db.Exec(stmt, word, translate)
	return err
}

func (s *Store) GetWords() ([]models.WordDB, error) {
	stmt := "SELECT word, translate FROM words"
	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	words := []models.WordDB{}

	for rows.Next() {
		word := models.WordDB{}
		if err := rows.Scan(&word.Word, &word.Translate); err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	return words, nil
}
