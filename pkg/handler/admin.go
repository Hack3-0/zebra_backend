package handler

import (
	"errors"
	"log"
	"net/http"
	"zebra/model"
	"zebra/utils"

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
	imageName, err := utils.CreateMenuItemImageImage(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
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
	menuItem.Image = imageName
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
