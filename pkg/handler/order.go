package handler

import (
	"log"
	"net/http"
	"zebra/model"
	"zebra/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) makeOrder(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	var reqOrder model.ReqOrder
	err = reqOrder.ParseRequest(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	log.Print(reqOrder)
	reqOrder.ID = id
	user, err := h.services.User.GetUserByID(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	log.Print(user)
	if user.Type == utils.TypeUser {
		org, err := h.services.Admin.GetOrgByID(reqOrder.OrganizationID)
		if err != nil {
			defaultErrorHandler(c, err)
			return
		}
		reqOrder.CashierID = org.Active
	}
	log.Printf("cash id is %v", reqOrder.CashierID)

	orderID, err := h.services.Order.GetNewOrderID()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	reqOrder.ID = orderID
	log.Printf("orderID is %v", orderID)
	err = h.services.Order.CreateOrder(reqOrder)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	coffeeNum := 0
	for _, item := range reqOrder.Items {
		if item.Category == "coffee" {
			coffeeNum++
		}
	}
	log.Print(reqOrder.UserID, reqOrder.Items)
	err = h.services.User.IncreaseCups(reqOrder.UserID, coffeeNum)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) changeOrderStatus(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	var ReqID model.ReqID
	err = ReqID.ParseRequest(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.services.Order.ChangeOrderStatus(ReqID.ID, id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) createFeedback(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	var ReqFeed model.ReqFeedback

	err = ReqFeed.ParseRequest(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	ReqFeed.UserID = id

	err = h.services.Order.CreateFeedback(ReqFeed)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(ReqFeed, c)
}

func (h *Handler) getOrders(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	cashier, err := h.services.GetCashByID(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	orders, err := h.services.Order.GetOrders(cashier.Organization.ID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(orders, c)
}
