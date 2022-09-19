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
	id, err := s.repo.GetNewNotificationID()
	if err != nil {
		return nil, err
	}
	data.ID = id
	return s.repo.CreateNotification(data)
}

func (s *LocalService) CreateNotificationAll(data *model.Notification, ids []int) error {
	for _, val := range ids {
		data.UserID = val
		_, err := s.CreateNotification(data)
		if err != nil {
			return err
		}
	}
	return nil
}
