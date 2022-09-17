package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Config struct {
	DBURI  string
	DBName string
}

const (
	collectionUser          = "users"
	collectionBoyman        = "boyman"
	collectionOrganizations = "organizations"
	collectionOrders        = "order"
	collectionMenu          = "menu"
	collectionMenuItems     = "menuItems"
)

func NewMongoDB(cfg Config) (*mongo.Database, *mongo.Client, error) {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DBURI))
	if err != nil {
		return nil, nil, err
	}

	database := client.Database(cfg.DBName)

	return database, client, nil
}

func GetMaxID(pipe []bson.M, WithIncrement bool, FirstID int, collection *mongo.Collection) (int, error) {
	var resID []*struct {
		ID int `json:"id" bson:"id"`
	}
	var newId int

	cursor, err := collection.Aggregate(context.TODO(), pipe)

	if err != nil { //Case aggregation with unwind gives error
		return 1, nil
	}

	if err = cursor.All(context.TODO(), &resID); err != nil {
		return 1, nil //Case aggregation with unwind gives error
	}

	if len(resID) > 0 {
		newId = resID[0].ID
		if WithIncrement {
			newId++
		}
	} else {
		newId = FirstID
	}

	return newId, nil
}

/*func GetUserData(studentsCollection *mongo.Collection, UserID int, Projectories ...string) (*model.User, error) {

	var users []model.User

	if len(Projectories) < 1 {
		return nil, errors.New("server error | GetUserData bad request")
	}

	proj := bson.M{}
	query := []bson.M{
		{"$match": bson.M{"id": UserID}},
	}

	r, _ := regexp.Compile("^joinData.")

	for _, val := range Projectories {
		if r.MatchString(val) {
			joinRequest := bson.M{
				"localField":   "id",
				"foreignField": "userId",
			}
			v := strings.Split(val, ".")
			if len(v) != 2 {
				return nil, errors.New("controller bad request")
			}
			joinRequest["as"] = val + "Array"

			joinRequest["from"] = v[1]

			query = append(query, bson.M{"$lookup": joinRequest})
			query = append(query, bson.M{"$addFields": bson.M{val: bson.M{"$arrayElemAt": []interface{}{"$" + val + "Array", 0}}}})
		}
		proj[val] = 1

	}

	if len(proj) > 0 {
		query = append(query, bson.M{"$project": proj})
	}
	// log.Print(query)

	cursor, err := studentsCollection.Aggregate(context.TODO(),
		query,
	)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return &users[0], err
}
*/
