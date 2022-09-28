package users

import (
	"delta-go/pkg/common/models"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h handler) GetUser(c *fiber.Ctx) error {
	email_param := c.Params("email")
	email := string(email_param)
	fmt.Println("Getting user with email", email)

	if email == "" || reflect.TypeOf(email).Kind() != reflect.String {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id")
	}

	var user models.User

	if result := h.DB.Where("email = ?", email).First(&user).Error; result == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, result.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}
