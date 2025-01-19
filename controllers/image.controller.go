package controllers

import (
	"bytes"
	"fmt"
	"image-processor-server/services"
	"io"

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

	fileBuffer := &bytes.Buffer{}
	if _, err := io.Copy(fileBuffer, multipartFile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to read image: " + err.Error(),
		})
	}

	head := fileBuffer.Bytes()[:512]

	if !services.IsValidImage(head) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid image type",
		})
	}

	fmt.Println("Tamaño del archivo:", len(fileBuffer.Bytes()))
	img, err := services.Rotate(fileBuffer.Bytes())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to process image: " + err.Error(),
		})
	}

	fmt.Println(img.ColorModel())

	return c.Status(fiber.StatusOK).SendString("Image upload successful")
}
