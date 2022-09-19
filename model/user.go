package model

import "time"

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
	Discount     float32            `json:"discount" bson:"discount"`
	Cups         int                `json:"cups" bson:"cups"`
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
	Title  string    `json:"title" bson:"title"`
	Text   string    `json:"text" bson:"text"`
	Time   time.Time `json:"time" bson:"time"`
}

type SendAllNotification struct {
	Title string `json:"title" bson:"title"`
	Text  string `json:"text" bson:"text"`
	Time  time.Time
}
