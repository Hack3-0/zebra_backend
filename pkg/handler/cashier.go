package handler

import (
	"errors"
	"log"
	"net/http"
	"zebra/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUpCash(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	var reqCash model.ReqCashRegistration
	err = reqCash.ParseRequest(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	org, err := h.services.Admin.GetOrgByID(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	reqCash.Organization = org
	log.Print(id)
	if err = h.services.Unauthed.CheckUsername(reqCash.Username); err == nil {
		defaultErrorHandler(c, errors.New("username is already taken"))
		return
	}
	newID, err := h.services.Unauthed.GetNewUserID()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	reqCash.ID = newID
	log.Print(reqCash)
	err = h.services.Cashier.CreateCash(reqCash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.Admin.AddCashier(id, newID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	res, err := h.services.Cashier.GetCashByUsername(reqCash.Username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendGeneral(res, c)
}

func (h *Handler) startSession(c *gin.Context) {
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

	org, err := h.services.Admin.GetOrgByID(cashier.Organization.ID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	if org.Active != 0 {
		defaultErrorHandler(c, errors.New("end previous session to start new one"))
		return
	}

	err = h.services.Cashier.StartSession(id, cashier.Organization.ID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) endSession(c *gin.Context) {
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
	err = h.services.Cashier.UpdateWorkHours(id, *cashier.StartTime)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.services.Cashier.EndSession(id, cashier.Organization.ID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) getCashiers(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	cashiers, err := h.services.GetCashiers(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(cashiers, c)
}
