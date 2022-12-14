package model

type MenuItem struct {
	ID          int     `json:"id" bson:"id"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Category    string  `json:"category" bson:"category"`
	Price       int     `json:"price" bson:"price"`
	Cost        int     `json:"cost" bson:"cost"`
	Discount    float32 `json:"discount" bson:"discount"`
	HasSuggar   bool    `json:"hasSugar" bson:"hasSugar"`
	Image       string  `json:"image" bson:"image"`
	Hide        bool    `json:"hide" bson:"hide"`
}

type StatMenu struct {
	ID       int    `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Revenue  int    `json:"revenue" bson:"revenue"`
	Cost     int    `json:"cost" bson:"cost"`
	Quantity int    `json:"quantity" bson:"quantity"`
}
