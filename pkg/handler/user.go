package handler

import (
	"errors"
	"strconv"
	"zebra/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) getUser(c *gin.Context) {
	keys := c.Request.URL.Query()["id"]
	id, err := strconv.Atoi(keys[0])
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	user, err := h.services.User.GetUserByID(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(user, c)
}

func (h *Handler) getCashier(c *gin.Context) {
	keys := c.Request.URL.Query()["id"]
	id, err := strconv.Atoi(keys[0])
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	user, err := h.services.Cashier.GetCashByID(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(user, c)
}

func (h *Handler) changeOrganization(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	var org model.ShortOrganization

	if err := c.ShouldBindWith(&org, binding.JSON); err != nil {
		defaultErrorHandler(c, errors.New("bad request | "+err.Error()))
	}
	err = h.services.User.ChangeOrganization(id, org)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) getUserInfo(c *gin.Context) {
	keys := c.Request.URL.Query()["id"]
	id, err := strconv.Atoi(keys[0])
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	user, err := h.services.User.GetUserByID(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	orders, err := h.services.User.GetUserOrders(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	type UserAllInfo struct {
		User   *model.User    `json:"user"`
		Orders []*model.Order `json:"orders"`
	}

	info := UserAllInfo{user, orders}

	sendGeneral(info, c)

}

func (h *Handler) getMenu(c *gin.Context) {
	menu, err := h.services.Menu.GetMenu()
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(menu, c)
}

func (h *Handler) getNotifications(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	notifications, err := h.services.User.GetNotifications(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(notifications, c)
}
