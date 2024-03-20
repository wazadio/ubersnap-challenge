package controllers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/wazadio/ubersnap/domain"
	"github.com/wazadio/ubersnap/helpers"
	"github.com/wazadio/ubersnap/services"
)

type imageController struct {
	imageService services.ImageService
}

type ImageController interface{}

func NewImageController(imageService services.ImageService) ImageController {
	return &imageController{
		imageService: imageService,
	}
}

func (c *imageController) ConvertPNGToJPG(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		helpers.Log(helpers.ERROR, err.Error())
		http.Error(w, "Failed to parse multipart form", http.StatusInternalServerError)

		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		helpers.Log(helpers.ERROR, err.Error())
		http.Error(w, "Failed to get image", http.StatusInternalServerError)

		return
	}

	if strings.ToLower(filepath.Ext(fileHeader.Filename)) != domain.PNG {
		http.Error(w, "image extension is not png", http.StatusBadRequest)

		return
	}

	jpgImage, err := c.imageService.ConvertPNGToJPG(file, fileHeader)
	if err != nil {
		helpers.Log(helpers.ERROR, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = w.Write(jpgImage)
	if err != nil {
		helpers.Log(helpers.ERROR, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	return
}
