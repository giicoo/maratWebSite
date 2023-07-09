package mongo_db

import (
	"context"

	"github.com/giicoo/maratWebSite/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Store) AddTest(test models.Test) error {
	_, err := s.collectionTests.InsertOne(context.TODO(), test)
	return err
}

func (s *Store) GetTestByName(name string) (models.Test, error) {
	filter := bson.M{"name": name}

	test := models.Test{}
	err := s.collectionTests.FindOne(context.TODO(), filter).Decode(&test)
	if err != nil {
		return test, err
	}
	return test, nil
}

func (s *Store) AddUserRes(res models.UserResult, test_name string) error {
	filter := bson.M{"name": test_name}

	_, err := s.collectionTests.UpdateMany(context.TODO(), filter, bson.D{{"$push", bson.D{{"usersresults", res}}}})
	return err
}

func (s *Store) GetTests() ([]*models.Test, error) {
	filter := bson.M{}
	tests := []*models.Test{}

	cur, err := s.collectionTests.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		test := models.Test{}
		err := cur.Decode(&test)
		if err != nil {
			return nil, err
		}
		tests = append(tests, &test)
	}

	return tests, nil
}
