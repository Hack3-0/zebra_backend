package handler

import (
	"errors"
	"zebra/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

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
	var menu *model.MenuItem
	if err := c.ShouldBindWith(&menu, binding.JSON); err != nil {
		defaultErrorHandler(c, errors.New("bad request | "+err.Error()))
		return
	}

	err := h.services.Menu.UpdateMenuItem(menu)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	menuItem, err := h.services.GetMenuItemByID(menu.ID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(menuItem, c)
}
