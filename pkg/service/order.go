package service

import (
	"time"
	"zebra/model"
	"zebra/pkg/fcmService"
	"zebra/pkg/repository"
)

type OrderService struct {
	repo         repository.Order
	pushService  fcmService.Push
	localService fcmService.Local
}

func NewOrderService(repo repository.Order, pushService fcmService.Push, localService fcmService.Local) *OrderService {
	return &OrderService{repo: repo, pushService: pushService, localService: localService}
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

func (s *OrderService) ChangeOrderStatus(id, cashID int) error {
	return s.repo.ChangeOrderStatus(id, cashID)
}

func (s *OrderService) CreateFeedback(req model.ReqFeedback) error {
	feedbackID, err := s.repo.GetNewFeedbackID()
	if err != nil {
		return err
	}

	req.ID = feedbackID
	return s.repo.CreateFeedback(req)
}

func (s *OrderService) GetOrders(id int) ([]*model.Order, error) {
	return s.repo.GetOrders(id)
}
