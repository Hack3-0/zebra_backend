package repository

import (
	"context"
	"zebra/model"
	"zebra/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type OrderMongo struct {
	db *mongo.Database
}

func NewOrderMongo(db *mongo.Database) *OrderMongo {
	return &OrderMongo{db: db}
}

func (s *OrderMongo) CreateOrder(data model.ReqOrder) error {
	col := s.db.Collection(collectionOrders)

	newOrder := &model.Order{}
	newOrder.ID = data.ID
	newOrder.UserID = data.UserID
	newOrder.CashierID = data.CashierID
	newOrder.OrganizationID = data.OrganizationID
	newOrder.Items = data.Items
	newOrder.Status = utils.Pending
	newOrder.Time = data.Time
	_, err := col.InsertOne(
		context.TODO(),
		newOrder,
	)

	return err
}

func (s *OrderMongo) GetOrderByID(id int) (*model.Order, error) {
	col := s.db.Collection(collectionOrders)

	var order *model.Order

	err := col.FindOne(
		context.TODO(),
		bson.M{"id": id},
	).Decode(&order)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderMongo) GetNewOrderID() (int, error) {
	collection := s.db.Collection(collectionOrders)
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
func (s *OrderMongo) GetNewFeedbackID() (int, error) {
	collection := s.db.Collection(collectionFeedback)
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

func (s *OrderMongo) ChangeOrderStatus(id int) error {
	col := s.db.Collection(collectionOrders)
	_, err := col.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.M{"$set": bson.M{"status": utils.Complete}},
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *OrderMongo) CreateFeedback(req model.ReqFeedback) error {
	col := s.db.Collection(collectionFeedback)

	newFeedback := &model.FeedBack{ID: req.ID, UserID: req.UserID, OrderID: req.OrderID, Text: req.Text, Rating: req.Rating}

	_, err := col.InsertOne(
		context.TODO(),
		newFeedback,
	)

	return err
}
