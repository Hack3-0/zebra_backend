package model

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ReqID struct
type ReqID struct {
	ID int `json:"id" bson:"id"`
}

func (p *ReqID) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.ID == 0 {
		return errors.New("bad request | id is required")
	}

	return nil
}

// ReqIDString struct
type ReqIDString struct {
	ID string `json:"id" bson:"id"`
}

func (p *ReqIDString) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.ID == "" {
		return errors.New("bad request | id is required")
	}

	return nil
}

type ReqUserRegistration struct {
	ID           int
	Token        string
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"password" bson:"password"`
	Name         string             `json:"name" bson:"name"`
	Preference   *Preference        `json:"preference" bson:"preference"`
	Organization *ShortOrganization `json:"organization" bson:"organization"`
	PushToken    string             `json:"pushToken" bson:"pushToken"`
	PhoneNumber  string             `json:"phoneNumber" bson:"pushToken"`
}

func (p *ReqUserRegistration) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.Username == "" || p.Password == "" {
		return errors.New("bad request | empty fields")
	}

	return nil
}

type ReqOrgRegistration struct {
	ID       int
	Token    string
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name"`
	Address  string `json:"address" bson:"address"`
}

func (p *ReqOrgRegistration) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.Username == "" || p.Password == "" {
		return errors.New("bad request | empty fields")
	}

	return nil
}

type ReqCashRegistration struct {
	ID           int
	Token        string
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"password" bson:"password"`
	Name         string             `json:"name" bson:"name"`
	Organization *ShortOrganization `json:"organization" bson:"organization"`
}

func (p *ReqCashRegistration) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.Username == "" || p.Password == "" {
		return errors.New("bad request | empty fields")
	}

	return nil
}

type ReqUserLogin struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	PushToken string `json:"pushToken"`
}

func (p *ReqUserLogin) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}

type ReqOrder struct {
	ID             int
	UserID         int         `json:"userID" bson:"userID"`
	CashierID      int         `json:"cashierID" bson:"cashierID"`
	OrganizationID int         `json:"organizationID" bson:"organizationID"`
	Items          []*MenuItem `json:"items" bson:"items"`
	Status         string
	Time           time.Time // can add start time and end time to check average completion time
}

func (p *ReqOrder) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.UserID == 0 || p.OrganizationID == 0 || len(p.Items) == 0 {
		log.Print(p.UserID, p.OrganizationID, p.Items)
		return errors.New("bad request | empty fields")
	}

	return nil
}

type ReqCheckUsername struct {
	Username string `json:"username" binding:"required"`

	BoymanData
}

func (p *ReqCheckUsername) ParseRequest(c *gin.Context) error {
	return c.BindJSON(&p)
}

// ReqName struct
type ReqName struct {
	Name string `json:"name" bson:"name"`
}

func (p *ReqName) ParseRequest(c *gin.Context) error {
	return c.ShouldBindWith(&p, binding.JSON)
}

type ReqTime struct {
	TimeStamp time.Time `json:"timeStamp" bson:"timeStamp"`
}

func (p *ReqTime) ParseRequest(c *gin.Context) error {
	return c.ShouldBindWith(&p, binding.JSON)
}

type ReqFeedback struct {
	ID      int
	UserID  int
	OrderID int    `json:"orderID" bson:"orderID"`
	Text    string `json:"text" bson:"text"`
	Rating  int    `json:"rating" bson:"rating"`
}

func (p *ReqFeedback) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.OrderID == 0 || p.Rating == 0 {
		return errors.New("bad request | empty fields")
	}

	return nil
}
