package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func CreateMenuItemImageImage(c *gin.Context) (string, error) {
	// Image handling
	c.Request.ParseMultipartForm(10 << 20)

	fileName := ""
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := c.Request.FormFile("image")
	if file != nil {
		defer file.Close()
	}

	if file != nil {
		defer file.Close()
	}

	locationImage, exists := os.LookupEnv("LocationMenuItems")

	if !exists {
		return "", errors.New("enviroment variable is not set")
	}

	locationBigImages := locationImage
	if err != nil {
		if err.Error() != "http: no such file" {

			return "", err
		}
	} else {
		// Create a temporary file within our temp-images directory that follows
		tempFile, err := ioutil.TempFile(locationBigImages, "upload-*.jpeg")
		if err != nil {
			return "", err
		}
		defer tempFile.Close()
		// read all of the contents of our uploaded file into a
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return "", err
		}
		tempFile.Write(fileBytes)

		fileName = filepath.Base(tempFile.Name())
	}
	c.Request.ParseMultipartForm(0)

	return fileName, nil
}
