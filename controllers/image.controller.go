package controllers

import (
	"bytes"
	"image-processor-server/services"
	"io"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HandleRotateImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	angleValue := c.Params("angle")
	angleInt, _ := strconv.Atoi(angleValue)

	multipartFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer multipartFile.Close()

	fileBuffer := &bytes.Buffer{}
	if _, err := io.Copy(fileBuffer, multipartFile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to read image " + err.Error(),
		})
	}

	head := fileBuffer.Bytes()[:512]

	if !services.IsValidImage(head) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid image type",
		})
	}

	rotatedImgBytes, err := services.Rotate(fileBuffer.Bytes(), float64(angleInt), "rotate")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to process image " + err.Error(),
		})
	}

	c.Set("Content-Type", "image/png")
	return c.Send(rotatedImgBytes)

}
