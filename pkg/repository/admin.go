package repository

import (
	"context"
	"log"
	"zebra/model"
	"zebra/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type AdminMongo struct {
	db *mongo.Database
}

func NewAdminMongo(db *mongo.Database) *AdminMongo {
	return &AdminMongo{db: db}
}

func (s *AdminMongo) CreateOrg(data model.ReqOrgRegistration) error {
	col := s.db.Collection(collectionUser)

	newOrg := &model.Organization{
		CashierID: []int{},
	}
	newOrg.ID = data.ID
	newOrg.Address = data.Address
	newOrg.Token = data.Token
	newOrg.Password = data.Password
	newOrg.Username = data.Username
	newOrg.Name = data.Name
	newOrg.Type = utils.TypeAdmin

	_, err := col.InsertOne(
		context.TODO(),
		newOrg,
	)

	return err
}

func (s *AdminMongo) GetOrgByID(id int) (*model.Organization, error) {
	col := s.db.Collection(collectionUser)

	var org *model.Organization

	err := col.FindOne(
		context.TODO(),
		bson.M{"id": id, "type": utils.TypeAdmin},
	).Decode(&org)
	log.Print(id)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (s *AdminMongo) GetNewOrgID() (int, error) {
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

func (s *AdminMongo) GetOrgByUsername(username string) (*model.Organization, error) {
	col := s.db.Collection(collectionUser)

	var org *model.Organization

	err := col.FindOne(
		context.TODO(),
		bson.M{"username": username, "type": utils.TypeAdmin},
	).Decode(&org)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (s *AdminMongo) GetOrganizations() ([]*model.Organization, error) {
	col := s.db.Collection(collectionUser)
	org := make([]*model.Organization, 0)
	cursor, err := col.Find(
		context.TODO(),
		bson.M{"type": utils.TypeAdmin},
	)

	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &org); err != nil {
		return nil, err
	}
	return org, nil
}
func (s *AdminMongo) AddCashier(id, newID int) error {
	col := s.db.Collection(collectionUser)
	_, err := col.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.M{"$push": bson.M{"cashierID": newID}},
	)
	if err != nil {
		return nil
	}

	return nil
}
