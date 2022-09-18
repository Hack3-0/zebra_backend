package handler

import (
	"strconv"
	"zebra/model"

	"github.com/gin-gonic/gin"
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

	var reqData struct {
		OrgID int `json:"organizationID" binding:"required"`
	}

	err = h.services.User.ChangeOrganization(id, reqData.OrgID)
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
