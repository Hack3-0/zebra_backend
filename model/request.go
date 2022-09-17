package model

import (
	"errors"

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

// ReqIDPage struct
type ReqIDPage struct {
	ID   int `json:"id" `
	Page int `json:"page"`
}

func (p *ReqIDPage) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.ID == 0 || p.Page == 0 {
		return errors.New("bad request | id & page is required")
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

// ReqFeedLikedUsers
type ReqFeedLikedUsers struct {
	ID     string `json:"id" bson:"id"`
	Page   int    `json:"page" bson:"page"`
	Search string `json:"search" bson:"search"`
}

func (p *ReqFeedLikedUsers) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.ID == "" {
		return errors.New("bad request | id is required")
	}

	if p.Page == 0 {
		return errors.New("bad request | page is required")
	}

	return nil
}

type ReqUserRegistration struct {
	ID                   int
	Token                string
	Username             string      `json:"username" bson:"username"`
	Password             string      `json:"password" bson:"password"`
	Type                 string      `json:"type" bson:"type"`
	Preference           *Preference `json:"preference" bson:"preference"`
	SelectedOrganization int         `json:"selectedOrganization" bson:"selectedOrganization"`
	PushToken            string      `json:"pushToken"`
	PhoneNumber          string      `json:"phoneNumber"`
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

type ReqUserLogin struct {
	Username  string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	PushToken string `json:"pushToken"`
	BoymanData
}

func (p *ReqUserLogin) ParseRequest(c *gin.Context) error {
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		return errors.New("bad request | " + err.Error())
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
