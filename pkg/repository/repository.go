package repository

import (
	"time"
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
	GetAllUserByUsername(Username string) (*model.AllUser, error)
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

type Menu interface {
	CreateMenuItem(data model.MenuItem) error
	GetMenuItemByID(id int) (*model.MenuItem, error)
	GetMenu() ([]*model.MenuItem, error)
	GetNewMenuItemID() (int, error)
	GetMenuCategory(category string) ([]*model.MenuItem, error)
	DeleteMenuItem(id int) error
	UpdateMenuItem(menu *model.MenuItem) error
}

type User interface {
	GetUserByID(id int) (*model.User, error)
	ChangeOrganization(id int, org model.ShortOrganization) error
	GetUserOrders(id int) ([]*model.Order, error)
}

type Admin interface {
	CreateOrg(data model.ReqOrgRegistration) error
	GetOrgByID(id int) (*model.Organization, error)
	GetOrgByUsername(username string) (*model.Organization, error)
	GetOrganizations() ([]*model.Organization, error)
	AddCashier(id, newID int) error
	CreateHeadAdmin(token, password, username, userType string) error
	GetRevenue(id int, timeStamp time.Time) (int, int, *model.MenuItem, error)
}
type Cashier interface {
	CreateCash(data model.ReqCashRegistration) error
	GetCashByID(id int) (*model.Cashier, error)
	GetCashByUsername(username string) (*model.Cashier, error)
	StartSession(id, orgID int) error
	UpdateWorkHours(id int, sessionDuration float32) error
	EndSession(id, orgID int) error
	GetCashiers(id int) ([]*model.Cashier, error)
	UpdateSession()
}
type LocalNotification interface {
	CreateNotification(data *model.Notification) (*model.Notification, error)
}
type Order interface {
	CreateOrder(data model.ReqOrder) error
	GetOrderByID(id int) (*model.Order, error)
	GetNewOrderID() (int, error)
	ChangeOrderStatus(id int) error
}

type Repository struct {
	Unauthed
	User
	Admin
	PushNotification
	Cashier
	Order
	Menu
	LocalNotification
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Unauthed: NewUnauthMongo(db),
		//Profile:           NewProfileMongo(db),
		PushNotification:  NewPushNotificationMongo(db),
		User:              NewUserMongo(db),
		Admin:             NewAdminMongo(db),
		Cashier:           NewCashierMongo(db),
		Order:             NewOrderMongo(db),
		Menu:              NewMenuMongo(db),
		LocalNotification: NewLocalNotificationMongo(db),
	}
}
