package model

import "time"

type Cashier struct {
	ID             int       `json:"id" bson:"id"`
	Token          string    `json:"token" bson:"token"`
	Username       string    `json:"username" bson:"username"`
	Password       string    `json:"password" bson:"password"`
	Name           string    `json:"name" bson:"name"`
	Type           string    `json:"type" bson:"type"`
	OrganizationID int       `json:"organizationID" bson:"organizationID"`
	StartTime      time.Time `json:"startTime" bson:"startTime"`
	HoursWorked    float32   `json:"hoursWorked" bson:"hoursWorked"`
}
