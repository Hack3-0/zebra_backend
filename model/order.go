package model

import "time"

type Order struct {
	ID             int         `json:"id" bson:"id"`
	UserID         int         `json:"userID" bson:"userID"`
	CashierID      int         `json:"cashierID" bson:"cashierID"`
	OrganizationID int         `json:"organizationID" bson:"organizationID"`
	Items          []*MenuItem `json:"menuItem" bson:"menuItem"`
	Status         string      `json:"status" bson:"status"`
	Time           time.Time   `json:"time" bson:"time"` // can add start time and end time to check average completion time
}
