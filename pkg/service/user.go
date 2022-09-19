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

func (s *UserService) ChangeOrganization(id int, org model.ShortOrganization) error {
	return s.repo.ChangeOrganization(id, org)
}

func (s *UserService) GetUserOrders(id int) ([]*model.Order, error) {
	orders, err := s.repo.GetUserOrders(id)
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		orders = []*model.Order{}
	}

	return orders, nil
}

func (s *UserService) IncreaseCups(id, coffeeNum int) error {
	return s.repo.IncreaseCups(id, coffeeNum)
}

func (s *UserService) GetNotifications(id int) ([]*model.Notification, error) {
	return s.repo.GetNotifications(id)
}

func (s *UserService) GetEveryUser() ([]int, []string, error) {
	return s.repo.GetEveryUser()
}
