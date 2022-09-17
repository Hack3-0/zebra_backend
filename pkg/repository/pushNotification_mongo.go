package repository

import (
	"context"

	"zebra/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type PushNotificationMongo struct {
	db *mongo.Database
}

func NewPushNotificationMongo(db *mongo.Database) *PushNotificationMongo {
	return &PushNotificationMongo{db: db}
}
func (s *PushNotificationMongo) GetPushToken(TakerID int) (string, error) {
	col := s.db.Collection(collectionUser)
	var user *model.User
	err := col.FindOne(
		context.TODO(),
		bson.M{"id": TakerID},
	).Decode(&user)
	if err != nil {
		return "", err
	}

	return user.PushToken, nil

}
