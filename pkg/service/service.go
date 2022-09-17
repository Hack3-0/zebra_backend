package service

import (
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
	CheckBoyman(Boyman, Timestamp string) error
	GetUserByPhone(Phone string) (*model.User, error)
	GetUserByUsername(Username string) (*model.User, error)
}

type Service struct {
	Unauthed
}

func NewService(repos *repository.Repository, fcmService *fcmService.FcmService) *Service {
	return &Service{
		Unauthed: NewUnauthService(repos.Unauthed, fcmService.Push),
	}
}
