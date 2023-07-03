package mongo_db

import (
	"context"

	"github.com/giicoo/maratWebSite/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Store) AddTest(test models.Test) error {
	collection := s.client.Database("maratDB").Collection("tests")
	_, err := collection.InsertOne(context.TODO(), test)
	return err
}

func (s *Store) GetTestByName(name string) (models.Test, error) {
	collection := s.client.Database("maratDB").Collection("tests")
	filter := bson.M{"name": name}

	test := models.Test{}
	err := collection.FindOne(context.TODO(), filter).Decode(&test)
	if err != nil {
		return test, err
	}
	return test, nil
}
