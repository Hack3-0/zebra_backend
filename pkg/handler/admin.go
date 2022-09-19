package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
	"zebra/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) signUpOrg(c *gin.Context) {

	var reqOrg model.ReqOrgRegistration
	err := reqOrg.ParseRequest(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	if err = h.services.Unauthed.CheckUsername(reqOrg.Username); err == nil {
		defaultErrorHandler(c, errors.New("username is already taken"))
		return
	}
	id, err := h.services.Unauthed.GetNewUserID()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	reqOrg.ID = id
	log.Printf("id %v", id)
	err = h.services.Admin.CreateOrg(reqOrg)
	log.Printf("req %v", reqOrg)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	res, err := h.services.Admin.GetOrgByUsername(reqOrg.Username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendGeneral(res, c)
}

func (h *Handler) getOrganizations(c *gin.Context) {
	organizations, err := h.services.Admin.GetOrganizations()

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(organizations, c)
}

func (h *Handler) createMenuItem(c *gin.Context) {
	var menuItem model.MenuItem
	if err := c.ShouldBindWith(&menuItem, binding.JSON); err != nil {
		defaultErrorHandler(c, errors.New("bad request | "+err.Error()))
		return
	}
	id, err := h.services.Menu.GetNewMenuItemID()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	menuItem.ID = id
	err = h.services.Menu.CreateMenuItem(menuItem)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	newMenuItem, err := h.services.Menu.GetMenuItemByID(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(newMenuItem, c)
}

func (h *Handler) getMenuCategory(c *gin.Context) {
	type ReqCategory struct {
		Category string `json:"category" bson:"category"`
	}
	var category ReqCategory
	if err := c.ShouldBindWith(&category, binding.JSON); err != nil {
		defaultErrorHandler(c, errors.New("bad request | "+err.Error()))
	}
	menu, err := h.services.Menu.GetMenuCategory(category.Category)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(menu, c)
}

func (h *Handler) getStatistics(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	var reqTime model.ReqTime
	err = reqTime.ParseRequest(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	log.Print(reqTime.TimeStamp.Date())
	if id == 0 {
		statistics, err := h.services.Admin.GetAllStatistics(reqTime.TimeStamp)
		if err != nil {
			defaultErrorHandler(c, err)
			return
		}
		sendGeneral(statistics, c)
	} else {
		statistics, err := h.services.Admin.GetStatistics(id, reqTime.TimeStamp)
		if err != nil {
			defaultErrorHandler(c, err)
			return
		}
		sendGeneral(statistics, c)
	}

}

func (h *Handler) getFeedback(c *gin.Context) {
	keys := c.Request.URL.Query()["id"]
	id, err := strconv.Atoi(keys[0])
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	log.Printf("Id %v", id)
	feedback, err := h.services.Admin.GetFeedback(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(feedback, c)
}

func (h *Handler) sendAll(c *gin.Context) {
	var notification model.SendAllNotification
	if err := c.ShouldBindWith(&notification, binding.JSON); err != nil {
		defaultErrorHandler(c, errors.New("bad request | "+err.Error()))
		return
	}
	realNotification := &model.Notification{Title: notification.Title, Text: notification.Text, Time: time.Now()}

	users, pushTokens, err := h.services.User.GetEveryUser()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	log.Print(users, pushTokens)

	h.localService.CreateNotificationAll(realNotification, users)
	go h.pushService.SendPushNotificationAll(pushTokens, notification.Text, notification.Title)
	sendSuccess(c)
}
