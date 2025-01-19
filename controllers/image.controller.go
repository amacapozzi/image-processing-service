package controllers

import (
	"image-processor-server/services"

	"github.com/gofiber/fiber/v2"
)

func HandleUploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	multipartFile, err := file.Open()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	defer multipartFile.Close()

	fileBuffer := make([]byte, 512)

	if _, err := multipartFile.Read(fileBuffer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to read image" + err.Error(),
		})
	}

	IS_VALID := services.IsValidImage(fileBuffer)

	if !IS_VALID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid image type",
		})
	}

	return c.Status(fiber.StatusOK).SendString("Image upload successful")
}
