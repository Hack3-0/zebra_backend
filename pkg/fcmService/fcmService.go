package fcmService

import (
	"zebra/pkg/repository"

	"github.com/appleboy/go-fcm"
	"go.mongodb.org/mongo-driver/mongo"
)

type Push interface {
	SendPushNotification(TakerID int, notificationType string) error
}

type Local interface {
}

type FcmService struct {
	Push
	Local
}

func NewFcmService(repos *repository.Repository, fcmClient *fcm.Client, db *mongo.Database) *FcmService {
	return &FcmService{
		Push:  NewPushService(repos, fcmClient),
		Local: NewLocalService(repos, fcmClient),
	}
}
