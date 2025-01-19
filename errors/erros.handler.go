package errors

import "github.com/gofiber/fiber/v2"

func HandleBadRequest(c *fiber.Ctx, err error) error {
	return c.Status(400).JSON(fiber.Map{
		"error": err.Error(),
	})
}
