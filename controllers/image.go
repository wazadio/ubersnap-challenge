package controllers

import (
	"net/http"

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

		w.
	}
}
