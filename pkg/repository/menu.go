package repository

import (
	"context"
	"errors"
	"zebra/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MenuMongo struct {
	db *mongo.Database
}

func NewMenuMongo(db *mongo.Database) *MenuMongo {
	return &MenuMongo{db: db}
}

func (s *MenuMongo) GetMenuItemByID(id int) (*model.MenuItem, error) {
	collection := s.db.Collection(collectionMenu)
	var menuItem *model.MenuItem
	err := collection.FindOne(
		context.TODO(),
		bson.M{"id": id},
	).Decode(&menuItem)

	if err != nil {
		return nil, err
	}

	return menuItem, nil
}

func (s *MenuMongo) GetMenu() ([]*model.MenuItem, error) {
	collection := s.db.Collection(collectionMenu)
	var menu []*model.MenuItem
	cursor, err := collection.Find(
		context.TODO(),
		bson.M{"hide": false},
	)

	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &menu); err != nil {
		return nil, err
	}

	return menu, nil
}

func (s *MenuMongo) CreateMenuItem(data model.MenuItem) error {
	collection := s.db.Collection(collectionMenu)
	_, err := collection.InsertOne(
		context.TODO(),
		data,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *MenuMongo) GetNewMenuItemID() (int, error) {
	collection := s.db.Collection(collectionMenu)
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

func (s *MenuMongo) GetMenuCategory(category string) ([]*model.MenuItem, error) {
	col := s.db.Collection(collectionMenu)
	var menu []*model.MenuItem
	cursor, err := col.Find(
		context.TODO(),
		bson.M{"category": category, "hide": false},
	)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &menu); err != nil {
		return nil, err
	}

	if len(menu) == 0 {
		menu = []*model.MenuItem{}
	}
	return menu, nil
}

func (s *MenuMongo) DeleteMenuItem(id int) error {
	col := s.db.Collection(collectionMenu)

	res, err := col.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.M{"$set": bson.M{"hide": true}},
	)
	if err != nil {
		return err
	}

	if res.UpsertedCount == 0 {
		return errors.New("not found")
	}
	return nil
}

func (s *MenuMongo) UpdateMenuItem(menu *model.MenuItem) error {
	col := s.db.Collection(collectionMenu)

	res, err := col.UpdateOne(
		context.TODO(),
		bson.M{"id": menu.ID},
		bson.M{"$set": menu},
	)
	if err != nil {
		return err
	}

	if res.UpsertedCount == 0 {
		return errors.New("not found")
	}
	return nil
}
