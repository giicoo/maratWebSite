package mongo_db

import (
	"context"

	"github.com/giicoo/maratWebSite/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Store) AddUser(user models.UserDB) error {
	collection := s.client.Database("maratDB").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func (s *Store) GetUser(login string) (models.UserDB, error) {
	collection := s.client.Database("maratDB").Collection("users")
	filter := bson.M{"login": login}

	user := models.UserDB{}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
