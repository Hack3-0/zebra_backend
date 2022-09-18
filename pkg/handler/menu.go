package handler

import (
	"log"
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
	var menuItem model.MenuItem
	id, err := strconv.Atoi(c.Request.FormValue("id"))
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	first, err := h.services.GetMenuItemByID(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	log.Print(first)
	menuItem.ID = id
	if c.Request.FormValue("image") != "" {
		imageName, err := utils.CreateMenuItemImageImage(c)
		if err != nil {
			defaultErrorHandler(c, err)
			return
		}
		menuItem.Image = imageName
	} else {
		menuItem.Image = first.Image
	}
	if c.Request.FormValue("name") == "" {
		menuItem.Name = first.Name
	}
	if c.Request.FormValue("description") == "" {
		menuItem.Description = first.Description
	}
	if c.Request.FormValue("category") == "" {
		menuItem.Category = first.Category
	}
	if c.Request.FormValue("price") == "" {
		menuItem.Price = first.Price
	} else {
		price, err := strconv.Atoi(c.Request.FormValue("price"))
		if err != nil {
			defaultErrorHandler(c, err)
			return
		}
		menuItem.Price = price

	}
	if c.Request.FormValue("discount") == "" {
		menuItem.Discount = first.Discount
	} else {
		discount, err := strconv.ParseFloat(c.Request.FormValue("discount"), 32)
		if err != nil {
			defaultErrorHandler(c, err)
			return
		}
		menuItem.Discount = float32(discount)
	}

	menuItem.HasSuggar = (c.Request.FormValue("hasSugar") == "true")

	log.Print(menuItem.ID)

	err = h.services.Menu.UpdateMenuItem(&menuItem)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	log.Print(menuItem)
	menu, err := h.services.GetMenuItemByID(menuItem.ID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	log.Print(menu)

	sendGeneral(menu, c)
}
