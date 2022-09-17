package handler

import "github.com/gin-gonic/gin"

func (h *Handler) getUser(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	user, err := h.services.User.GetUserByID(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(user, c)
}

func (h *Handler) changeOrganization(c *gin.Context) {
	id, err := getUserId(c)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	var reqData struct {
		OrgID int `json:"selectedOrganization" binding:"required"`
	}

	err = h.services.User.ChangeOrganization(id, reqData.OrgID)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendSuccess(c)
}
