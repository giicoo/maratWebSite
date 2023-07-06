package mongo_db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	client          *mongo.Client
	collectionWords *mongo.Collection
	collectionTests *mongo.Collection
	collectionUsers *mongo.Collection
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) InitDB() error {
	// create connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// ping DB
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	// set client
	s.client = client

	// init collection
	s.collectionWords = s.client.Database("maratDB").Collection("words")
	s.collectionTests = s.client.Database("maratDB").Collection("tests")
	s.collectionUsers = s.client.Database("maratDB").Collection("users")
	return nil
}
