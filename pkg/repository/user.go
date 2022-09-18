package repository

import (
	"context"
	"errors"
	"zebra/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserMongo struct {
	db *mongo.Database
}

func NewUserMongo(db *mongo.Database) *UserMongo {
	return &UserMongo{db: db}
}

func (s *UserMongo) GetUserByID(id int) (*model.User, error) {
	collection := s.db.Collection(collectionUser)
	var user *model.User
	err := collection.FindOne(
		context.TODO(),
		bson.M{"id": id},
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserMongo) ChangeOrganization(id int, org model.Organization) error {
	collection := s.db.Collection(collectionUser)

	res, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.M{"$set": bson.M{"organization": org}},
	)

	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("not found")
	}

	return nil
}

func (s *UserMongo) GetUserOrders(id int) ([]*model.Order, error) {
	col := s.db.Collection(collectionOrders)
	var orders []*model.Order

	cursor, err := col.Find(
		context.TODO(),
		bson.M{"userID": id},
	)

	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &orders); err != nil {
		return nil, err
	}

	return orders, nil
}
