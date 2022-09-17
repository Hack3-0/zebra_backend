package service

import (
	"time"
	"zebra/model"
	"zebra/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(data model.ReqOrder) error {
	data.Time = time.Now()
	err := s.repo.CreateOrder(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) GetOrderByID(id int) (*model.Order, error) {
	return s.repo.GetOrderByID(id)
}

func (s *OrderService) GetNewOrderID() (int, error) {
	return s.repo.GetNewOrderID()
}
