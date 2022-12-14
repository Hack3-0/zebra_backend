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

type ShortOrganization struct {
	ID      int    `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
}

type Statistics struct {
	ID          int         `json:"id" bson:"id"`
	Address     string      `json:"address" bson:"address"`
	Revenue     int         `json:"revenue" bson:"revenue"`
	Margin      float32     `json:"margin" bson:"margin"`
	ProductStat []*StatMenu `json:"productStat" bson:"productStat"`
}

type StatFeedback struct {
	Organization *ShortOrganization `json:"organization" bson:"organization"`
	Rating       float32            `json:"rating" bson:"rating"`
	Feedback     []*FeedBack        `json:"feedback" bson:"feedback"`
}
