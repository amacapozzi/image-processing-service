package controllers

import (
	"bytes"
	"image-processor-server/errors"
	"image-processor-server/services"
	"io"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HandleRotateImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		errors.HandleBadRequest(c, err)
	}

	angleValue := c.Params("angle")
	angleInt, _ := strconv.Atoi(angleValue)

	multipartFile, err := file.Open()
	if err != nil {
		errors.HandleBadRequest(c, err)
	}
	defer multipartFile.Close()

	fileBuffer := &bytes.Buffer{}
	if _, err := io.Copy(fileBuffer, multipartFile); err != nil {
		errors.HandleBadRequest(c, err)
	}

	head := fileBuffer.Bytes()[:512]

	if !services.IsValidImage(head) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid image type",
		})
	}

	rotatedImgBytes, err := services.Rotate(fileBuffer.Bytes(), float64(angleInt), "rotate")
	if err != nil {
		errors.HandleBadRequest(c, err)
	}

	c.Set("Content-Type", "image/png")
	return c.Send(rotatedImgBytes)

}

func HandleGrayScale(c *fiber.Ctx) error {

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	formData, err := c.MultipartForm()

	if err != nil {
		errors.HandleBadRequest(c, err)
	}

	multipartFile, err := file.Open()
	if err != nil {
		errors.HandleBadRequest(c, err)
	}
	defer multipartFile.Close()

	var imgBuffer bytes.Buffer

	if _, err := io.Copy(&imgBuffer, multipartFile); err != nil {
		errors.HandleBadRequest(c, err)
	}

	widthStr, heightStr := formData.Value["width"][0], formData.Value["height"][0]

	width, _ := strconv.Atoi(widthStr)
	height, _ := strconv.Atoi(heightStr)

	resizedImage, err := services.Resize(imgBuffer.Bytes(), width, height)

	if err != nil {
		errors.HandleBadRequest(c, err)
	}

	c.Set("Content-Type", "image/png")
	return c.Send(resizedImage)

}
