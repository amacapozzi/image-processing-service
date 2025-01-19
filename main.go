package main

import (
	"image-processor-server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupImageRoutes(app)

	app.Listen(":3000")
}
