package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"zebra/model"
	"zebra/utils"

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
	menuItem.Name = c.Request.FormValue("name")
	menuItem.Description = c.Request.FormValue("description")
	menuItem.Category = c.Request.FormValue("category")
	menuItem.Price, err = strconv.Atoi(c.Request.FormValue("price"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	discount, err := strconv.ParseFloat(c.Request.FormValue("discount"), 32)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	menuItem.Discount = float32(discount)
	menuItem.HasSuggar = (c.Request.FormValue("hasSugar") == "true")

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
