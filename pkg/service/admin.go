package service

import (
	"log"
	"time"
	"zebra/model"
	"zebra/pkg/repository"

	"github.com/dgrijalva/jwt-go"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) CreateOrg(data model.ReqOrgRegistration) error {
	token, err := s.GenerateToken(data.ID)
	if err != nil {
		return err
	}
	data.Token = token
	log.Print(data.Token)
	err = s.repo.CreateOrg(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *AdminService) GetOrgByID(id int) (*model.Organization, error) {
	return s.repo.GetOrgByID(id)
}

func (s *AdminService) GenerateToken(UserID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},
		UserID,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AdminService) GetOrgByUsername(username string) (*model.Organization, error) {
	return s.repo.GetOrgByUsername(username)
}

func (s *AdminService) GetOrganizations() ([]*model.Organization, error) {
	organizations, err := s.repo.GetOrganizations()
	if err != nil {
		return nil, err
	}

	if len(organizations) == 0 {
		organizations = []*model.Organization{}
	}

	return organizations, err
}

func (s *AdminService) AddCashier(id, newID int) error {
	return s.repo.AddCashier(id, newID)
}

func (s *AdminService) GetAllStatistics(timeStamp time.Time) ([]*model.Statistics, error) {
	admins, err := s.repo.GetOrganizations()
	if err != nil {
		return nil, err
	}
	stat := make([]*model.Statistics, 0)
	for _, item := range admins {
		s, err := s.GetStatistics((item.ID), timeStamp)
		if err != nil {
			return nil, err
		}
		stat = append(stat, s)
	}
	return stat, nil
}

func (s *AdminService) GetStatistics(id int, timeStamp time.Time) (*model.Statistics, error) {
	var res *model.Statistics
	revenue, cost, popular, err := s.repo.GetRevenue(id, timeStamp)
	if err != nil {
		return nil, err
	}
	log.Print(revenue, cost, popular)
	res.ID = id
	res.Revenue = revenue
	res.Margin = float32(revenue-cost) / float32(revenue)
	res.PopularProduct = popular
	admin, err := s.repo.GetOrgByID(id)
	if err != nil {
		return nil, err
	}
	res.Address = admin.Address

	return res, nil
}
