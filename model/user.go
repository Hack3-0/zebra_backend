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
	ID           int                `json:"id" bson:"id"`
	Token        string             `json:"token" bson:"token"`
	PushToken    string             `json:"pushToken" bson:"pushToken"`
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"password" bson:"password"`
	Name         string             `json:"name" bson:"name"`
	PhoneNumber  string             `json:"phoneNumber" bson:"phoneNumber"`
	Type         string             `json:"type" bson:"type"`
	Preference   *Preference        `json:"preference" bson:"preference"`
	Organization *ShortOrganization `json:"organization" bson:"organization"`
}

type Preference struct {
	Sugar      bool        `json:"sugar" bson:"sugar"`
	CoffeeType []*MenuItem `json:"coffeeType" bson:"coffeeType"`
	MilkType   []*MenuItem `json:"milkType" bson:"milkType"`
}

type UserResponse struct {
	ID           int                `json:"id" bson:"id"`
	Token        string             `json:"token" bson:"token"`
	PushToken    string             `json:"pushToken" bson:"pushToken"`
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"password" bson:"password"`
	Name         string             `json:"name" bson:"name"`
	PhoneNumber  string             `json:"phoneNumber" bson:"phoneNumber"`
	Type         string             `json:"type" bson:"type"`
	Preference   *Preference        `json:"preference" bson:"preference"`
	Organization *ShortOrganization `json:"organization" bson:"organization"`
}

type AllUser struct {
	ID       int    `json:"id" bson:"id"`
	Token    string `json:"token" bson:"token"`
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
	Type     string `json:"type" bson:"type"`
}

type Notification struct {
	ID     int       `json:"id" bson:"id"`
	UserID int       `json:"userID" bson:"userID"`
	Title  string    `json:"string" bson:"string"`
	Text   string    `json:"Text" bson:"Text"`
	Time   time.Time `json:"time" bson:"time"`
}
