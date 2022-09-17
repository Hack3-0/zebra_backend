package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"zebra/model"
	"zebra/pkg/fcmService"
	"zebra/pkg/repository"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt             = "hjqrhjqw124617ajfhajs"
	signingKey       = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL         = 12 * time.Hour
	boymanLocalToken = "ba2228a00d21e19c23e4f210a5b8a300"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"userId"`
}

type UnauthedService struct {
	repo        repository.Unauthed
	pushService fcmService.Push
}

func NewUnauthService(repo repository.Unauthed, pushService fcmService.Push) *UnauthedService {
	return &UnauthedService{repo: repo, pushService: pushService}
}

func (s *UnauthedService) CreateUser(user model.ReqUserRegistration) error {
	id, err := s.GetNewUserID()

	if err != nil {
		return err
	}
	user.ID = id
	log.Print(user.ID)
	user.Token, err = s.GenerateToken(user.ID)
	if err != nil {
		return err
	}
	log.Print(user.Token)

	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UnauthedService) SetPushToken(email string, pushToken string) error {
	return s.repo.SetPushToken(email, pushToken)
}

func (s *UnauthedService) CheckUsername(username string) error {
	return s.repo.CheckUsername(username)
}

func (s *UnauthedService) GenerateToken(UserID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},
		UserID,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *UnauthedService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func (s *UnauthedService) GetNewUserID() (int, error) {
	return s.repo.GetNewUserID()
}

func (s *UnauthedService) CheckBoyman(Boyman, Timestamp string) error {
	if Boyman == "test" && Timestamp == "test" { //used for development and testing
		return nil
	}

	hasher := md5.New()
	hasher.Write([]byte(Timestamp + boymanLocalToken))
	b := hex.EncodeToString(hasher.Sum(nil))

	if b != Boyman {
		return errors.New("incorrect boyman")
	}

	err := s.repo.BoymanExists(Timestamp)

	if err == nil {
		s.repo.InsertBoyman(Timestamp, Boyman)
	}

	return err
}

func (s *UnauthedService) GetUserByPhone(Phone string) (*model.User, error) {
	return s.repo.GetUserByPhone(Phone)
}

func (s *UnauthedService) GetUserByUsername(Username string) (*model.User, error) {
	return s.repo.GetUserByUsername(Username)
}
