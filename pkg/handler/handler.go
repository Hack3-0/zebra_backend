package handler

import (
	"zebra/pkg/service"

	"github.com/gin-contrib/cors"
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
		unauthed.Use(cors.Default())
		unauthed.POST("/signup", h.signUp)
		unauthed.POST("/signin", h.signIn)
		unauthed.POST("/getOrganizations", h.getOrganizations)
		unauthed.GET("/getMenu", h.getMenu)
		unauthed.POST("/getMenuCategory", h.getMenuCategory)

	}
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NjM0NDEyMzcsInVzZXJJZCI6M30.2nL_PBWJ7MXZia6Q0e9xdMuysk3ijkY5J1yL_FKRgZE

	authed := router.Group("/authed", h.userIdentity)
	{
		authed.POST("/startSession", h.startSession)
		authed.POST("/endSession", h.endSession)
		authed.POST("/makeOrder", h.makeOrder)
		authed.POST("/changeOrderStatus", h.changeOrderStatus)
		authed.GET("/getUser", h.getUser)
		authed.POST("/changeOrganization", h.changeOrganization)
		authed.GET("/getUserInfo", h.getUserInfo)
		admin := authed.Group("/admin", h.adminIdentity)
		{
			admin.POST("/signup", h.signUpCash)
			admin.POST("/getCashiers", h.getCashiers)
			admin.POST("/getStatistics", h.getStatistics)
			headAdmin := admin.Group("/headAdmin", h.headAdminIdentity)
			{
				headAdmin.POST("/uploadImage", h.uploadImage)
				headAdmin.POST("/deleteMenuItem", h.deleteMenuItem)
				headAdmin.POST("/updateMenuItem", h.updateMenuItem)
				headAdmin.POST("/signup", h.signUpOrg)
				headAdmin.POST("/createMenuItem", h.createMenuItem)
			}
		}
	}

	return router
}
