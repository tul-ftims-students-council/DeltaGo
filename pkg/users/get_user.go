package users

import (
	"delta-go/pkg/common/models"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUser(c *fiber.Ctx) error {
	email := c.Params("email")
	fmt.Println("Getting user with email", email)

	if email == "" || reflect.TypeOf(email).Kind() != reflect.String {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid email type")
	}

	var user models.User

	// TODO: Change responses to use HanldeError from models
	if result := h.DB.Where("Email == ?", email).First(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}
