package mongo_db

import (
	"context"

	"github.com/giicoo/maratWebSite/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Store) AddUser(user models.User) error {
	_, err := s.collectionUsers.InsertOne(context.TODO(), user)
	return err
}

func (s *Store) GetUserByLogin(login string) (models.User, error) {
	filter := bson.M{"login": login}

	user := models.User{}

	err := s.collectionUsers.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
