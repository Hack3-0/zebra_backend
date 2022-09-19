package handler

import (
	"errors"
	"zebra/model"
	"zebra/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) getMenuItem(c *gin.Context) {
	var id model.ReqID
	err := id.ParseRequest(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	item, err := h.services.Menu.GetMenuItem(id.ID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(item, c)
}
func (h *Handler) deleteMenuItem(c *gin.Context) {
	var id model.ReqID
	err := id.ParseRequest(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	err = h.services.Menu.DeleteMenuItem(id.ID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}

func (h *Handler) updateMenuItem(c *gin.Context) {
	var menuItem model.MenuItem
	if err := c.ShouldBindWith(&menuItem, binding.JSON); err != nil {
		defaultErrorHandler(c, errors.New("bad request | "+err.Error()))
	}

	err := h.services.Menu.UpdateMenuItem(&menuItem)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	menu, err := h.services.GetMenuItemByID(menuItem.ID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(menu, c)
}

func (h *Handler) uploadImage(c *gin.Context) {
	imageName, err := utils.CreateMenuItemImageImage(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(imageName, c)
}
