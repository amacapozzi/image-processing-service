package services

import (
	"bytes"
	"image/color"
	"image/png"
	"math"

	"github.com/disintegration/imaging"
	"github.com/h2non/filetype"
)

func IsValidImage(imgBytes []byte) bool {
	return filetype.IsImage(imgBytes)
}

func Rotate(imgBytes []byte, angle float64, direction string) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, err
	}

	angle = normalizeAngle(angle)

	if direction == "counterclockwise" {
		angle = -angle
	}

	rotatedImage := imaging.Rotate(img, angle, color.Transparent)

	var buf bytes.Buffer
	if err := png.Encode(&buf, rotatedImage); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func normalizeAngle(angle float64) float64 {
	normalized := math.Mod(angle, 360)
	if normalized < 0 {
		normalized += 360
	}
	return normalized
}

func Resize(imgBytes []byte, width int, height int) ([]byte, error) {
	image, err := imaging.Decode(bytes.NewReader(imgBytes), imaging.AutoOrientation(true))

	if err != nil {
		return nil, err
	}

	NRGBA := imaging.Resize(image, width, height, imaging.Lanczos)

	var buf bytes.Buffer

	if err := imaging.Encode(&buf, NRGBA, imaging.PNG); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
