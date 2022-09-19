package service

import (
	"log"
	"sort"
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
	var res model.Statistics
	revenue, cost, popular, err := s.repo.GetRevenue(id, timeStamp)
	if err != nil {
		return nil, err
	}
	log.Print(revenue, cost, popular)
	res.ID = id
	res.Revenue = revenue
	if revenue == 0 {
		res.Margin = 0
	} else {
		res.Margin = float32(revenue-cost) / float32(revenue) * float32(100)
	}
	if len(popular) == 0 {
		res.ProductStat = []*model.StatMenu{}
	} else {
		res.ProductStat = popular

	}
	sort.Slice(res.ProductStat, func(i, j int) bool {
		return res.ProductStat[i].Quantity > res.ProductStat[j].Quantity
	})
	admin, err := s.repo.GetOrgByID(id)
	if err != nil {
		return nil, err
	}
	res.Address = admin.Address

	return &res, nil
}

func (s *AdminService) GetFeedback(id int) (*model.StatFeedback, error) {
	var feedbacks *model.StatFeedback
	org, err := s.repo.GetOrgByID(id)
	if err != nil {
		return nil, err
	}
	feedbacks.Organization = &model.ShortOrganization{ID: org.ID, Address: org.Address, Name: org.Name}
	rating, feedback, err := s.repo.GetFeedback(id)
	if err != nil {
		return nil, err
	}
	if len(feedback) == 0 {
		feedback = []*model.FeedBack{}
	}
	feedbacks.Rating = rating
	feedbacks.Feedback = feedback

	return feedbacks, nil
}
