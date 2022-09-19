package fcmService

import (
	"log"

	"zebra/pkg/repository"

	"github.com/appleboy/go-fcm"
	"github.com/sirupsen/logrus"
)

type PushService struct {
	fcmClient *fcm.Client
	repo      repository.PushNotification
}

func NewPushService(repo repository.PushNotification, fcmClient *fcm.Client) *PushService {
	return &PushService{repo: repo, fcmClient: fcmClient}
}

func (s *PushService) SendPushNotification(TakerID int, text, title string) error {

	pushToken, err := s.repo.GetPushToken(TakerID)
	if err != nil {
		return err
	}
	if pushToken == "" {
		logrus.Print("empty pushToken")
		return nil
	}
	msg := &fcm.Message{
		To: pushToken,
		Notification: &fcm.Notification{
			Title:       title,
			Body:        text,
			ClickAction: "FLUTTER_NOTIFICATION_CLICK",
		},
	}
	// Send the message and receive the response without retries.
	_, err = s.fcmClient.Send(msg)
	if err != nil {
		log.Print("ERROR: ", err.Error())
		return nil
	}
	return nil
}
