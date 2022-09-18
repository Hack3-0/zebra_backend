package service

import (
	"zebra/model"
	"zebra/pkg/repository"
)

type MenuService struct {
	repo repository.Menu
}

func NewMenuService(repo repository.Menu) *MenuService {
	return &MenuService{repo: repo}
}

func (s *MenuService) CreateMenuItem(data model.MenuItem) error {
	err := s.repo.CreateMenuItem(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *MenuService) GetMenuItemByID(id int) (*model.MenuItem, error) {
	return s.repo.GetMenuItemByID(id)
}

func (s *MenuService) GetMenu() ([]*model.MenuItem, error) {
	menu, err := s.repo.GetMenu()
	if err != nil {
		return nil, err
	}

	if len(menu) == 0 {
		menu = []*model.MenuItem{}
	}

	return menu, err
}

func (s *MenuService) GetNewMenuItemID() (int, error) {
	return s.repo.GetNewMenuItemID()
}

func (s *MenuService) GetMenuCategory(category string) ([]*model.MenuItem, error) {
	return s.repo.GetMenuCategory(category)
}

func (s *MenuService) DeleteMenuItem(id int) error {
	return s.repo.DeleteMenuItem(id)
}

func (s *MenuService) UpdateMenuItem(menu *model.MenuItem) error {
	return s.repo.UpdateMenuItem(menu)
}
