package service

import (
	"time"
	"zebra/model"
	"zebra/pkg/fcmService"
	"zebra/pkg/repository"
)

type Unauthed interface {
	CreateUser(model.ReqUserRegistration) error
	SetPushToken(string, string) error
	CheckUsername(string) error
	GenerateToken(int) (string, error)
	GetNewUserID() (int, error)
	ParseToken(token string) (int, error)
	GetUserByPhone(Phone string) (*model.User, error)
	GetUserByUsername(Username string) (*model.User, error)
	GetAllUserByUsername(Username string) (*model.AllUser, error)
	//CheckCredentials(Username, Password string) error
}

type Admin interface {
	CreateOrg(data model.ReqOrgRegistration) error
	GetOrgByID(id int) (*model.Organization, error)
	GetOrgByUsername(username string) (*model.Organization, error)
	GetOrganizations() ([]*model.Organization, error)
	AddCashier(id, newID int) error
	GetStatistics(id int, timeStamp time.Time) (*model.Statistics, error)
	GetAllStatistics(timeStamp time.Time) ([]*model.Statistics, error)
}

type Cashier interface {
	CreateCash(data model.ReqCashRegistration) error
	GetCashByID(id int) (*model.Cashier, error)
	GetCashByUsername(username string) (*model.Cashier, error)
	StartSession(id, orgID int) error
	UpdateWorkHours(id int, startTime time.Time) error
	EndSession(id, orgID int) error
	GetCashiers(id int) ([]*model.Cashier, error)
}

type User interface {
	GetUserByID(id int) (*model.User, error)
	ChangeOrganization(id int, org model.ShortOrganization) error
	GetUserOrders(id int) ([]*model.Order, error)
	IncreaseCups(id, coffeeNum int) error
}

type Menu interface {
	CreateMenuItem(data model.MenuItem) error
	GetMenuItemByID(id int) (*model.MenuItem, error)
	GetMenu() ([]*model.MenuItem, error)
	GetNewMenuItemID() (int, error)
	GetMenuCategory(category string) ([]*model.MenuItem, error)
	DeleteMenuItem(id int) error
	UpdateMenuItem(menu *model.MenuItem) error
	GetMenuItem(id int) (*model.MenuItem, error)
}

type Order interface {
	CreateOrder(data model.ReqOrder) error
	GetOrderByID(id int) (*model.Order, error)
	GetNewOrderID() (int, error)
	CreateFeedback(req model.ReqFeedback) error
	ChangeOrderStatus(id int) error
	GetOrders(id int) ([]*model.Order, error)
}
type Service struct {
	Unauthed
	User
	Admin
	Cashier
	Order
	Menu
}

func NewService(repos *repository.Repository, fcmService *fcmService.FcmService) *Service {
	return &Service{
		Unauthed: NewUnauthService(repos.Unauthed, fcmService.Push),
		User:     NewUserService(repos.User),
		Admin:    NewAdminService(repos.Admin),
		Cashier:  NewCashierService(repos.Cashier),
		Order:    NewOrderService(repos.Order),
		Menu:     NewMenuService(repos.Menu),
	}
}
