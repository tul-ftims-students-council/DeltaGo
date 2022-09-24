package users

import (
	"delta-go/pkg/common/models"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Getting user with id", id)

	if id == "" || reflect.TypeOf(id).Kind() != reflect.Int {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id")
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}
