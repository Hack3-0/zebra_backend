package repository

import (
	"context"
	"zebra/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type LocalNotificationMongo struct {
	db *mongo.Database
}

func NewLocalNotificationMongo(db *mongo.Database) *LocalNotificationMongo {
	return &LocalNotificationMongo{db: db}
}

func (s *LocalNotificationMongo) CreateNotification(data *model.Notification) (*model.Notification, error) {
	col := s.db.Collection(collectionNotification)

	_, err := col.InsertOne(
		context.TODO(),
		data,
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *LocalNotificationMongo) GetNotifications(id int) ([]*model.Notification, error) {
	col := s.db.Collection(collectionNotification)
	var notifications []*model.Notification
	cursor, err := col.Find(
		context.TODO(),
		bson.M{
			"userID": id,
		},
	)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &notifications); err != nil {
		return nil, err
	}

	return notifications, nil
}
