package fcmService

import (
	"time"
	"zebra/model"
	"zebra/pkg/repository"

	"github.com/appleboy/go-fcm"
)

type LocalService struct {
	fcmClient *fcm.Client
	repo      repository.LocalNotification
}

func NewLocalService(repo repository.LocalNotification, fcmClient *fcm.Client) *LocalService {
	return &LocalService{repo: repo, fcmClient: fcmClient}
}

func (s *LocalService) CreateNotification(data *model.Notification) (*model.Notification, error) {
	data.Time = time.Now()
	return s.repo.CreateNotification(data)
}
