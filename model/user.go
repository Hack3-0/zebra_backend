package model

import "time"

type Notifications struct {
	Id     int       `json:"id" bson:"id"`
	UserId int       `json:"userId" bson:"userId"`
	Type   string    `json:"type" bson:"type"`
	Title  string    `json:"title,omitempty" bson:"title,omitempty"`
	Text   string    `json:"text,omitempty" bson:"text,omitempty"`
	Time   time.Time `json:"time" bson:"time"`
	Seen   bool      `json:"seen" bson:"seen"`
}

type User struct {
	ID             int         `json:"id" bson:"id"`
	Token          string      `json:"token" bson:"token"`
	PushToken      string      `json:"pushToken" bson:"pushToken"`
	Username       string      `json:"username" bson:"username"`
	Password       string      `json:"password" bson:"password"`
	Name           string      `json:"name" bson:"name"`
	PhoneNumber    string      `json:"phoneNumber" bson:"phoneNumber"`
	Type           string      `json:"type" bson:"type"`
	Preference     *Preference `json:"preference" bson:"preference"`
	OrganizationID int         `json:"organizationID" bson:"organizationID"`
}

type Preference struct {
	Sugar      bool  `json:"sugar" bson:"sugar"`
	CoffeeType []int `json:"coffeeType" bson:"coffeeType"`
	Milk       []int `json:"milk" bson:"milk"`
}
