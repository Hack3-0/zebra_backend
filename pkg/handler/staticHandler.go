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

	router.GET("/:locationType/:imageType/:fileName", h.RouterImageHandler)

	return router
}

func (h *StaticHandler) RouterImageHandler(c *gin.Context) {

	filename := c.Param("fileName")
	imageType := c.Param("imageType")
	locationType := c.Param("locationType")

	if filename == "" || imageType == "" || locationType == "" {
		defaultErrorHandler(c, errors.New("bad request"))
		return
	}

	if locationType != "customExercises" && locationType != "profileImage" && locationType != "LocationWorkoutUploadImage" && locationType != "LocationMeasurementsUploadImage" {
		defaultErrorHandler(c, errors.New("bad request"))
		return
	}

	if imageType != "bigImage" && imageType != "smallImage" {
		defaultErrorHandler(c, errors.New("bad request"))
		return
	}

	var locationPrefix string

	switch locationType {
	case "customExercises":
		locationPrefix = os.Getenv("LocationCustomExercises")

	case "profileImage":
		locationPrefix = os.Getenv("LocationProfileImage")

	case "LocationWorkoutUploadImage":
		locationPrefix = os.Getenv("LocationWorkoutUploadImage")

	case "LocationMeasurementsUploadImage":
		locationPrefix = os.Getenv("LocationMeasurementsUploadImage")

	}
	locationPrefix += imageType + "/"

	logrus.Print(locationPrefix + filename)

	c.File(locationPrefix + filename)
}
