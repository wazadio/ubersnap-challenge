package services

import (
	"bytes"
	"image"
	"image/jpeg"
	"mime/multipart"

	"github.com/wazadio/ubersnap/helpers"
)

type imageService struct{}

type ImageService interface {
	ConvertPNGToJPG(file multipart.File, fileHeader multipart.FileHeader) (jpgImage []byte, err error)
}

func NewImageService() ImageService {
	return &imageService{}
}

func (s *imageService) ConvertPNGToJPG(file multipart.File, fileHeader multipart.FileHeader) (jpgImage []byte, err error) {
	defer file.Close()

	// Decode the file data into an image
	img, _, err := image.Decode(file)
	if err != nil {
		helpers.Log(helpers.ERROR, err.Error())

		return
	}

	// Convert the image to a JPEG format in memory
	jpegData := bytes.NewBuffer(nil)
	err = jpeg.Encode(jpegData, img, nil)
	if err != nil {
		helpers.Log(helpers.ERROR, err.Error())

		return
	}

	return jpegData.Bytes(), err

	// imgInfo, err := png.DecodeConfig(file)
	// if err != nil {
	// 	helpers.Log(helpers.ERROR, err.Error())

	// 	return
	// }

	// Create a buffer to store the file contents
	// var buffer bytes.Buffer

	// // Copy the file contents into the buffer
	// _, err = io.Copy(&buffer, file)
	// if err != nil {
	// 	helpers.Log(helpers.ERROR, err.Error())

	//     return
	// }

	// // Convert buffer to byte slice
	// fileBytes := buffer.Bytes()

	// mat, err := gocv.NewMatFromBytes(imgInfo.Height, imgInfo.Width, gocv.MatTypeCV8UC3, fileBytes)
	// if err != nil {
	// 	helpers.Log(helpers.ERROR, err.Error())

	// 	return
	// }

	// ok := gocv.IMWrite(fileHeader.Filename[:len(fileHeader.Filename)-3], mat)
}
