package model

const (
	statusSuccess = 1000
)

type BoymanData struct {
	Boyman    string `json:"boyman" bson:"boyman" binding:"required"`
	Timestamp string `json:"timestamp" bson:"timestamp" binding:"required"`
}

// DefaultPage
type DefaultPage struct {
	TotalPages  int64       `json:"totalPages" bson:"totalPages"`
	CurrentPage int         `json:"currentPage" bson:"currentPage"`
	Data        interface{} `json:"data" bson:"data"`
}

// DefaultResponse used in every response
type DefaultResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

// SuccessResponse returns success
func SuccessResponse() *DefaultResponse {
	dr := &DefaultResponse{StatusCode: statusSuccess}
	return dr
}
