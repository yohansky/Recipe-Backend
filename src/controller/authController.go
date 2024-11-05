package controller

import (
	"backend-recipe/src/config"
	"backend-recipe/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["Password"] != data["Passwordconfirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "Password do not match",
		})
	}

	user := models.User{
		Name:  data["Name"],
		Email: data["Email"],
		Phone: data["Phone"],
	}

	user.SetPassword(data["Password"])

	config.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	config.DB.Where("email = ?", data["Email"]).First(&user)

	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "Email nor found",
		})
	}

	if err := user.ComparePassword(data["Password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "incorrect Password",
		})
	}

	item := map[string]string{
		"Email":   data["Email"],
		"Message": "Login Success",
		"Id":      strconv.Itoa(int(user.Id)),
	}

	return c.JSON(item)
}

// func Logout(c *fiber.Ctx) error {

// }
