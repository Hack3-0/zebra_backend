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
