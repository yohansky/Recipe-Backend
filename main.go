package main

import (
	"backend-recipe/src/config"
	"backend-recipe/src/helper"
	"backend-recipe/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migrate()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.Router(app)

	app.Listen(":8080")
}
