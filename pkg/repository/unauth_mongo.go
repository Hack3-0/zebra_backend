package repository

import (
	"context"
	"errors"

	"zebra/model"
	"zebra/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UnauthMongo struct {
	db *mongo.Database
}

func NewUnauthMongo(db *mongo.Database) *UnauthMongo {
	return &UnauthMongo{db: db}
}

func (r *UnauthMongo) GetAllUserByUsername(Username string) (*model.AllUser, error) {
	collection := r.db.Collection(collectionUser)
	var user *model.AllUser
	err := collection.FindOne(
		context.TODO(),
		bson.M{"username": Username},
	).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UnauthMongo) CreateUser(user model.ReqUserRegistration) error {
	usersCollection := r.db.Collection(collectionUser)

	newUser := &model.User{
		PhoneNumber: "",
	}
	newUser.ID = user.ID
	newUser.Token = user.Token
	newUser.PushToken = user.PushToken
	newUser.Username = user.Username
	newUser.Preference = user.Preference
	newUser.Type = utils.TypeUser
	newUser.Organization = user.Organization
	newUser.Password = user.Password
	newUser.PhoneNumber = user.PhoneNumber
	newUser.Name = user.Name

	// TODO: if googleAvatar is empty, set default random image

	_, err := usersCollection.InsertOne(
		context.TODO(),
		newUser,
	)

	return err
}

func (r *UnauthMongo) SetPushToken(email string, pushToken string) error {
	col := r.db.Collection("users")
	_, err := col.UpdateOne(
		context.TODO(),
		bson.M{"username": email},
		bson.M{"$set": bson.M{
			"pushToken": pushToken,
		}})

	return err
}

func (r *UnauthMongo) CheckUsername(username string) error {
	col := r.db.Collection("users")
	res := col.FindOne(
		context.TODO(),
		bson.M{"username": username})

	return res.Err()
}

func (r *UnauthMongo) GetNewUserID() (int, error) {
	collection := r.db.Collection("users")
	var resID []*model.ReqID
	var newId int

	cursor, err := collection.Aggregate(context.TODO(), []bson.M{
		{"$group": bson.M{"_id": nil, "id": bson.M{"$max": "$id"}}}},
	)

	if err != nil { //Case aggregation with unwind gives error
		return 1, nil
	}

	if err = cursor.All(context.TODO(), &resID); err != nil {
		return 1, nil //Case aggregation with unwind gives error
	}

	if len(resID) > 0 {
		newId = resID[0].ID
		newId++
	} else {
		newId = 1
	}

	return newId, nil
}

func (r *UnauthMongo) BoymanExists(timestamp string) error {
	boymansCollection := r.db.Collection(collectionBoyman)

	var boymanSaved model.BoymanData

	err := boymansCollection.FindOne(
		context.TODO(),
		bson.M{"timestamp": timestamp},
	).Decode(&boymanSaved)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

	return errors.New("incorrect boyman | boyman expired")
}

func (r *UnauthMongo) InsertBoyman(timestamp, boyman string) {
	boymansCollection := r.db.Collection(collectionBoyman)

	boymanSaved := &model.BoymanData{
		Timestamp: timestamp,
		Boyman:    boyman,
	}

	boymansCollection.InsertOne(
		context.TODO(),
		boymanSaved,
	)
}

func (r *UnauthMongo) GetUserByPhone(Phone string) (*model.User, error) {
	collection := r.db.Collection(collectionUser)

	var user *model.User

	err := collection.FindOne(
		context.TODO(),
		bson.M{"phoneNumber": Phone},
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UnauthMongo) GetUserByUsername(Username string) (*model.User, error) {
	collection := r.db.Collection(collectionUser)
	var user *model.User
	err := collection.FindOne(
		context.TODO(),
		bson.M{"username": Username},
	).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
