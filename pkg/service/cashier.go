package service

import (
	"log"
	"time"
	"zebra/model"
	"zebra/pkg/repository"

	"github.com/dgrijalva/jwt-go"
)

type CashierService struct {
	repo repository.Cashier
}

func NewCashierService(repo repository.Cashier) *CashierService {
	return &CashierService{repo: repo}
}

func (s *CashierService) CreateCash(data model.ReqCashRegistration) error {
	token, err := s.GenerateToken(data.ID)
	if err != nil {
		return err
	}
	data.Token = token
	log.Print(data.Token)
	err = s.repo.CreateCash(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *CashierService) GetCashByID(id int) (*model.Cashier, error) {
	return s.repo.GetCashByID(id)
}

func (s *CashierService) GenerateToken(UserID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},
		UserID,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *CashierService) GetCashByUsername(username string) (*model.Cashier, error) {
	return s.repo.GetCashByUsername(username)
}

func (s *CashierService) StartSession(id, orgID int) error {
	return s.repo.StartSession(id, orgID)
}

func (s *CashierService) UpdateWorkHours(id int, startTime time.Time) error {
	now := time.Now()
	sessionDuration := float32(now.Hour()-startTime.Hour()) + float32(now.Minute()-startTime.Minute())/60
	return s.repo.UpdateWorkHours(id, sessionDuration)
}

func (s *CashierService) EndSession(id, orgID int) error {
	return s.repo.EndSession(id, orgID)
}

func (s *CashierService) GetCashiers(id int) ([]*model.Cashier, error) {
	cashiers, err := s.repo.GetCashiers(id)
	if err != nil {
		return nil, err
	}
	if len(cashiers) == 0 {
		cashiers = []*model.Cashier{}
	}
	return cashiers, nil
}
