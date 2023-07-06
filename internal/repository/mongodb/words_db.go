package mongo_db

import (
	"context"

	"github.com/giicoo/maratWebSite/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Store) AddWord(word models.Word) error {
	_, err := s.collectionWords.InsertOne(context.TODO(), word)
	return err
}

func (s *Store) GetWords() ([]*models.Word, error) {

	filter := bson.M{}
	words := []*models.Word{}

	cur, err := s.collectionWords.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		word := models.Word{}
		err := cur.Decode(&word)
		if err != nil {
			return nil, err
		}
		words = append(words, &word)
	}

	return words, nil
}

func (s *Store) GetWordsByNames(words []*models.Word) ([]*models.Word, error) {
	elm_fil := []bson.M{}
	for _, item := range words {
		elm_fil = append(elm_fil, bson.M{"word": item.Word})
	}
	filter := bson.M{"$or": elm_fil}

	answers := []*models.Word{}

	cur, err := s.collectionWords.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		word := models.Word{}
		err := cur.Decode(&word)
		if err != nil {
			return nil, err
		}
		answers = append(answers, &word)
	}

	return answers, nil
}
