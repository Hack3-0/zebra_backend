package handler

import (
	"errors"
	"log"
	"net/http"

	"zebra/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {

	var reqUser model.ReqUserRegistration
	err := reqUser.ParseRequest(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	log.Print(reqUser)
	if err = h.services.Unauthed.CheckUsername(reqUser.Username); err == nil {
		defaultErrorHandler(c, errors.New("username is already taken"))
		return
	}
	log.Print(reqUser.Username)
	if reqUser.PhoneNumber != "" {
		user, err := h.services.Unauthed.GetUserByPhone(reqUser.PhoneNumber)
		if err != nil && err.Error() != "mongo: no documents in result" {
			defaultErrorHandler(c, err)
			return
		}
		log.Print(reqUser.PhoneNumber)

		if user != nil {
			defaultErrorHandler(c, errors.New("phone number is already taken"))
			return
		}
	}

	err = h.services.Unauthed.CreateUser(reqUser)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	res, err := h.services.GetAllUserByUsername(reqUser.Username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendGeneral(res, c)
}

func (h *Handler) signIn(c *gin.Context) {
	var reqData model.ReqUserLogin

	err := reqData.ParseRequest(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	user, err := h.services.GetUserByUsername(reqData.Username)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	if user.Password != reqData.Password {
		defaultErrorHandler(c, errors.New("username or password is wrong"))
		return
	}

	h.services.SetPushToken(reqData.Username, reqData.PushToken)

	AllUser := model.AllUser{
		ID:       user.ID,
		Token:    user.Token,
		Username: user.Username,
		Name:     user.Name,
		Type:     user.Type,
	}

	sendGeneral(AllUser, c)
}

/*
func (h *Handler) checkUsername(c *gin.Context) {
	var reqData model.ReqCheckUsername

	err := reqData.ParseRequest(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.services.Unauthed.CheckUsername(reqData.Username)

	if err == nil {
		defaultErrorHandler(c, errors.New("username is already taken"))
		return
	}

	sendSuccess(c)
}*/
