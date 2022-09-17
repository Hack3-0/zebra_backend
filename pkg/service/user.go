package service

import (
	"zebra/model"
	"zebra/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id int) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) ChangeOrganization(id, orgID int) error {
	return s.repo.ChangeOrganization(id, orgID)
}
