package handler

import (
	"zebra/pkg/fcmService"
	"zebra/pkg/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     *service.Service
	pushService  fcmService.Push
	localService fcmService.Local
}

func NewHandler(services *service.Service, pushService fcmService.Push, localService fcmService.Local) *Handler {
	return &Handler{services: services, pushService: pushService, localService: localService}
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
		authed.GET("/getCashier", h.getCashier)
		authed.POST("/changeOrganization", h.changeOrganization)
		authed.GET("/getUserInfo", h.getUserInfo)
		authed.POST("/createFeedback", h.createFeedback)
		authed.POST("/getOrders", h.getOrders)
		authed.POST("/getNotifications", h.getNotifications)
		admin := authed.Group("/admin", h.adminIdentity)
		{
			admin.POST("/signup", h.signUpCash)
			admin.POST("/getCashiers", h.getCashiers)
			admin.POST("/getStatistics", h.getStatistics)
			admin.POST("/getFeedback", h.getFeedback)
			headAdmin := admin.Group("/headAdmin", h.headAdminIdentity)
			{
				headAdmin.POST("/sendAll", h.sendAll)
				headAdmin.POST("/uploadImage", h.uploadImage)
				headAdmin.POST("/deleteMenuItem", h.deleteMenuItem)
				headAdmin.POST("/updateMenuItem", h.updateMenuItem)
				headAdmin.POST("/signup", h.signUpOrg)
				headAdmin.POST("/createMenuItem", h.createMenuItem)
				headAdmin.POST("/getMenuItem", h.getMenuItem)
			}
		}
	}

	return router
}
