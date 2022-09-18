package model

type Organization struct {
	ID        int    `json:"id" bson:"id"`
	Token     string `json:"token" bson:"token"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Name      string `json:"name" bson:"name"`
	Address   string `json:"address" bson:"address"`
	CashierID []int  `json:"cashierID" bson:"cashierID"`
	Type      string `json:"type" bson:"type"`
	Active    int    `json:"active" bson:"active"`
}

type Statistics struct {
	ID             int       `json:"id" bson:"id"`
	Address        string    `json:"address" bson:"address"`
	Revenue        int       `json:"revenue" bson:"revenue"`
	Margin         float32   `json:"margin" bson:"margin"`
	PopularProduct *MenuItem `json:"menuItem" bson:"menuItem"`
}
