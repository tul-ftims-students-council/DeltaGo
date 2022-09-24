package users

import (
	"delta-go/pkg/common/models"
	"delta-go/pkg/common/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h handler) Register(c *fiber.Ctx) error {
	fmt.Println("Registering user")
	body := models.User{}

	if err := c.BodyParser(&body); err != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	user.Name = body.Name
	user.Surname = body.Surname
	user.Email = body.Email
	user.PhoneNumber = body.PhoneNumber

	if result := h.DB.Where("Email = ?", body.Email).First(&user); result.Error == nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, "Użytkownik z takim mailem już istnieje")
	}

	if result := h.DB.Create(&user); result.Error != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, result.Error.Error())
	}

	return utils.HandleResponse(c, fiber.StatusOK, "Użytkownik został zarejestrowany")
}
