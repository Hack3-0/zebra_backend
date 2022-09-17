package handler

import (
	"errors"
	"net/http"
	"zebra/model"

	"github.com/gin-gonic/gin"
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
	err = h.services.Admin.CreateOrg(reqOrg)
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
