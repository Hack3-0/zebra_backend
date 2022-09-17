package handler

import (
	"errors"
	"os"
	"zebra/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	unauthed := router.Group("/unauthed")
	{
		unauthed.POST("/signup", h.signUp)
		unauthed.POST("/signin", h.signIn)
		unauthed.POST("/getOrganizations", h.getOrganizations)

	}
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NjM0NDEyMzcsInVzZXJJZCI6M30.2nL_PBWJ7MXZia6Q0e9xdMuysk3ijkY5J1yL_FKRgZE

	authed := router.Group("/authed", h.userIdentity)
	{
		authed.POST("/startSession", h.startSession)
		authed.POST("/endSession", h.endSession)
		authed.POST("/makeOrder", h.makeOrder)
		//authed.POST("/changeOrderStatus", h.changeOrderStatus)
		authed.GET("/getUser", h.getUser)
		authed.POST("/changeOrganization", h.changeOrganization)
		admin := authed.Group("/admin", h.adminIdentity)
		{
			admin.POST("/signup", h.signUpCash)
			headAdmin := admin.Group("/headAdmin", h.headAdminIdentity)
			{
				headAdmin.POST("/signup", h.signUpOrg)
			}
		}
	}

	return router
}

// mediaHandler used to handle topic thread
func (h *Handler) mediaHandler(c *gin.Context) {
	filename := c.Param("fileName")
	if filename == "" {
		defaultErrorHandler(c, errors.New("bad request"))
		return
	}

	locationMedia := os.Getenv("LocationMedia")

	c.File(locationMedia + filename)
}
