package handler

import (
	"errors"
	"os"

	"zebra/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type StaticHandler struct {
	services *service.Service
}

func NewStaticHandler(services *service.Service) *StaticHandler {
	return &StaticHandler{services: services}
}

func (h *StaticHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("getQr/:fileName", h.RouterImageHandler)
	router.GET("menu/:fileName", h.MenuImageHandler)
	return router
}

func (h *StaticHandler) MenuImageHandler(c *gin.Context) {

	filename := c.Param("fileName")

	if filename == "" {
		defaultErrorHandler(c, errors.New("bad request"))
		return
	}

	locationPrefix := os.Getenv("LocationMenuItems")

	logrus.Print(locationPrefix + filename)

	c.File(locationPrefix + filename)
}

func (h *StaticHandler) RouterImageHandler(c *gin.Context) {

	filename := c.Param("fileName")

	if filename == "" {
		defaultErrorHandler(c, errors.New("bad request"))
		return
	}

	var locationPrefix string

	locationPrefix = os.Getenv("LocationQr")

	logrus.Print(locationPrefix + filename)

	c.File(locationPrefix + filename)
}
