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
	err = h.services.User.IncreaseCups(reqOrder.UserID, coffeeNum)

	sendSuccess(c)
}

func (h *Handler) changeOrderStatus(c *gin.Context) {
	var ReqID model.ReqID
	err := ReqID.ParseRequest(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.services.Order.ChangeOrderStatus(ReqID.ID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}
