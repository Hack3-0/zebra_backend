package handler

import (
	"errors"
	"os"
	"time"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://3.70.198.111"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	unauthed := router.Group("/unauthed")
	{
		unauthed.POST("/signup", h.signUp)
		unauthed.POST("/signin", h.signIn)
		unauthed.POST("/getOrganizations", h.getOrganizations)
		unauthed.GET("/getMenu", h.getMenu)
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
			headAdmin := admin.Group("/headAdmin", h.headAdminIdentity)
			{
				headAdmin.POST("/signup", h.signUpOrg)
				headAdmin.POST("/createMenuItem", h.createMenuItem)
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
