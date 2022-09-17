package repository

import (
	"zebra/model"

	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Unauthed interface {
	CreateUser(user model.ReqUserRegistration) error
	SetPushToken(string, string) error
	CheckUsername(string) error
	GetNewUserID() (int, error)
	BoymanExists(string) error
	InsertBoyman(Timestamp, Boyman string)
	GetUserByPhone(Phone string) (*model.User, error)
	GetUserByUsername(Username string) (*model.User, error)
}

type Profile interface {
	//GetOwn(UserID int) (*model.OwnProfileData, error)
}

type PushNotification interface {
	GetPushToken(TakerID int) (string, error)
	/*GetFriendRequestData(UserID int) (pushToken, username, language string, err error)
	CheckSettingsRepo(UserID int, notificationType string) (bool, error)
	GetShortUser(TakerID, SenderID int) (*model.ShortUserSuggest, error)*/
}

type Repository struct {
	Unauthed
	//Profile
	PushNotification
	//LocalNotification
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Unauthed: NewUnauthMongo(db),
		//Profile:           NewProfileMongo(db),
		PushNotification: NewPushNotificationMongo(db),
	}
}
