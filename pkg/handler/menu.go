package handler

import (
	"strconv"
	"zebra/model"
	"zebra/utils"

	"github.com/gin-gonic/gin"
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
	var menuItem *model.MenuItem

	if c.Request.FormValue("image") != "" {
		imageName, err := utils.CreateMenuItemImageImage(c)
		if err != nil {
			defaultErrorHandler(c, err)
			return
		}
		menuItem.Image = imageName
	}
	menuItem.Name = c.Request.FormValue("name")
	menuItem.Description = c.Request.FormValue("description")
	menuItem.Category = c.Request.FormValue("category")
	price, err := strconv.Atoi(c.Request.FormValue("price"))
	menuItem.Price = price
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

	err = h.services.Menu.UpdateMenuItem(menuItem)
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
