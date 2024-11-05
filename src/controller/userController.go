package controller

import (
	"backend-recipe/src/config"
	"backend-recipe/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.User{}, page))
}
