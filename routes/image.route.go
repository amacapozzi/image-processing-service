package routes

import (
	"image-processor-server/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupImageRoutes(app *fiber.App) {
	app.Post("/image/rotate/:angle", controllers.HandleRotateImage)
	app.Post("/image/resize", controllers.HandleGrayScale)
}
