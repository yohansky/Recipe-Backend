package routes

import (
	"backend-recipe/src/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/register", controller.Register)
}
