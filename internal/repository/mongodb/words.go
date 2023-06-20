package mongo_db

import (
	"context"
	"fmt"

	"github.com/giicoo/maratWebSite/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Store) AddWord(word models.WordDB) error {
	collection := s.client.Database("maratDB").Collection("words")
	_, err := collection.InsertOne(context.TODO(), word)
	return err
}

func (s *Store) GetWords() ([]*models.WordDB, error) {
	collection := s.client.Database("maratDB").Collection("words")

	filter := bson.M{}
	words := []*models.WordDB{}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		word := models.WordDB{}
		err := cur.Decode(&word)
		if err != nil {
			return nil, err
		}
		words = append(words, &word)
	}
	fmt.Println(words[0].Word)
	return words, nil
}

func (s *Store) GetWordsByNames(words []*models.WordDB) ([]*models.WordDB, error) {
	// collection := s.client.Database("maratDB").Collection("words")

	filter := bson.M{}
	for _, item := range words {
		filter["word"] = item.Word
	}
	fmt.Println(filter)

	answers := []*models.WordDB{}

	return answers, nil
}
