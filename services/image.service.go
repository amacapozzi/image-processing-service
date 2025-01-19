package services

import (
	"bytes"
	"image"
	"image/png"

	"github.com/h2non/filetype"
)

func IsValidImage(imgBytes []byte) bool {
	return filetype.IsImage(imgBytes)
}

func Rotate(imgBytes []byte) (image image.Image, err error) {

	img, err := png.Decode(bytes.NewReader(imgBytes))

	return img, err
}
