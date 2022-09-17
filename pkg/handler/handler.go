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
		unauthed.GET("/hell", h.hello)
		//unauthed.POST("/signin", h.signIn)
		//unauthed.POST("/checkUsername", h.checkUsername)
	}
	/*
		authed := router.Group("/authed", h.userIdentity)
		{

		}*/
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
