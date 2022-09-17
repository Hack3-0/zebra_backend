package repository

import (
	"context"
	"log"
	"time"
	"zebra/model"
	"zebra/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type CashierMongo struct {
	db *mongo.Database
}

func NewCashierMongo(db *mongo.Database) *CashierMongo {
	return &CashierMongo{db: db}
}

func (s *CashierMongo) CreateCash(data model.ReqCashRegistration) error {
	col := s.db.Collection(collectionUser)

	newOrg := &model.Cashier{}
	newOrg.ID = data.ID
	newOrg.Token = data.Token
	newOrg.Password = data.Password
	newOrg.Username = data.Username
	newOrg.Name = data.Name
	newOrg.Type = utils.TypeCashier
	newOrg.OrganizationID = data.OrganizationID
	_, err := col.InsertOne(
		context.TODO(),
		newOrg,
	)

	return err
}

func (s *CashierMongo) GetCashByID(id int) (*model.Cashier, error) {
	col := s.db.Collection(collectionUser)

	var cash *model.Cashier

	err := col.FindOne(
		context.TODO(),
		bson.M{"id": id, "type": utils.TypeCashier},
	).Decode(&cash)
	log.Print(id)
	if err != nil {
		return nil, err
	}

	return cash, nil
}

func (s *CashierMongo) GetNewCashID() (int, error) {
	collection := s.db.Collection(collectionUser)
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

func (s *CashierMongo) GetCashByUsername(username string) (*model.Cashier, error) {
	col := s.db.Collection(collectionUser)

	var cash *model.Cashier

	err := col.FindOne(
		context.TODO(),
		bson.M{"username": username, "type": utils.TypeCashier},
	).Decode(&cash)
	if err != nil {
		return nil, err
	}

	return cash, nil
}

func (s *CashierMongo) StartSession(id, orgID int) error {
	col := s.db.Collection(collectionUser)

	_, err := col.UpdateOne(
		context.TODO(),
		bson.M{"id": orgID},
		bson.M{"$set": bson.M{"active": id}},
	)
	if err != nil {
		return err
	}
	_, err = col.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.M{"$set": bson.M{"startTime": time.Now()}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *CashierMongo) UpdateWorkHours(id int, sessionDuration float32) error {
	col := s.db.Collection(collectionUser)

	_, err := col.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.M{"$inc": bson.M{"hoursWorked": sessionDuration}},
	)
	if err != nil {
		return err
	}

	return nil
}
